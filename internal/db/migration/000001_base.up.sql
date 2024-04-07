
CREATE TABLE "Tag" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" Text UNIQUE,
  "description" Text
);

CREATE TABLE "Team" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" Text,
  "avatarUrl" Text,
  "subDomain" Text,
  "theme" Text,
  "canComment" Boolean,
  "canShare" Boolean,
  "inviteRequired" Boolean,
  "defaultUserRole" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "url" Text
);

CREATE TABLE "Workspace" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" Text,
  "domain" Text,
  "isPublic" Boolean,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "url" Text
);

CREATE TABLE "Profile" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "userId" Text UNIQUE,
  "fullName" Text,
  "age" Int,
  "address" Text,
  "avatarUrl" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp
);

CREATE TABLE "User" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "email" Text UNIQUE,
  "username" Text UNIQUE,
  "password" Text,
  "firstName" Text,
  "lastName" Text,
  "isAdmin" Boolean DEFAULT false,
  "isActive" Boolean,
  "isViewer" Boolean,
  "language" Text,
  "lastActiveAt" Timestamp,
  "roles" Text[] DEFAULT ARRAY[]::text[],
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp
);

CREATE TABLE "Document" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" Text,
  "text" Text,
  "emoji" Text,
  "isPublic" Boolean,
  "isFullWidth" Boolean,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "authorId" Text,
  "teamId" Text,
  "workspaceId" Text,
  "collectionId" Text
);

CREATE TABLE "Collection" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" Text,
  "description" Text,
  "icon" Text,
  "isSaving" Boolean,
  "canShare" Boolean,
  "childCollectionIds" Text,
  "downloadPermission" Text[] DEFAULT ARRAY[]::text[],
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "workspaceId" Text,
  "parentCollectionId" Text,
  "ownerUserId" Text
);
CREATE TABLE "Member" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "userId" Text,
  "collectionId" Text,
  "role" Text[] DEFAULT ARRAY[]::text[],
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp
);

CREATE TABLE "ActionEvent" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "action" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "actor" Text,
  "assignee" Text,
  "assigner" Text,
  "documentId" Text,
  "collectionId" Text
);

CREATE TABLE "Comment" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "content" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "userId" Text,
  "documentId" Text,
  "parentCommentId" Text
);

CREATE TABLE "Position" (
  "id" Text PRIMARY KEY DEFAULT uuid_generate_v4(),
  "line" Int,
  "col" Int,
  "toLine" Text,
  "toCol" Text,
  "content" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "commentsId" Text
);

CREATE TABLE "User_Team" (
  "user_id" Text,
  "team_id" Text
);

CREATE TABLE "User_Workspace" (
  "user_id" Text,
  "workspace_id" Text
);

CREATE TABLE "User_Document" (
  "user_id" Text,
  "document_id" Text
);

CREATE TABLE "Team_Workspace" (
  "team_id" Text,
  "workspace_id" Text
);

CREATE TABLE "Team_Document" (
  "team_id" Text,
  "document_id" Text
);

CREATE TABLE "Workspace_Document" (
  "workspace_id" Text,
  "document_id" Text
);

CREATE TABLE "Document_Collection" (
  "document_id" Text,
  "collection_id" Text
);

CREATE TABLE "Collection_Member" (
  "collection_id" Text,
  "member_id" Text
);

CREATE TABLE "Document_Tag" (
  "document_id" Text,
  "tag_id" Text
);

CREATE TABLE "Collection_Tag" (
  "folder_id" Text,
  "tag_id" Text
);

COMMENT ON COLUMN "Team"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Workspace"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Profile"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "User"."lastActiveAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "User"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Document"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Collection"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Member"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Comment"."updatedAt" IS 'Automatically updated by the database';

COMMENT ON COLUMN "Position"."updatedAt" IS 'Automatically updated by the database';

ALTER TABLE "Profile" ADD FOREIGN KEY ("userId") REFERENCES "User" ("id");

ALTER TABLE "Document" ADD FOREIGN KEY ("teamId") REFERENCES "Team" ("id");

ALTER TABLE "Document" ADD FOREIGN KEY ("workspaceId") REFERENCES "Workspace" ("id");

ALTER TABLE "Document" ADD FOREIGN KEY ("collectionId") REFERENCES "Collection" ("id");

ALTER TABLE "Collection" ADD FOREIGN KEY ("workspaceId") REFERENCES "Workspace" ("id");

ALTER TABLE "Collection" ADD FOREIGN KEY ("parentCollectionId") REFERENCES "Collection" ("id");

ALTER TABLE "Collection" ADD FOREIGN KEY ("ownerUserId") REFERENCES "User" ("id");

ALTER TABLE "Member" ADD FOREIGN KEY ("userId") REFERENCES "User" ("id");

ALTER TABLE "Member" ADD FOREIGN KEY ("collectionId") REFERENCES "Collection" ("id");

ALTER TABLE "ActionEvent" ADD FOREIGN KEY ("documentId") REFERENCES "Document" ("id");

ALTER TABLE "ActionEvent" ADD FOREIGN KEY ("collectionId") REFERENCES "Collection" ("id");

ALTER TABLE "Comment" ADD FOREIGN KEY ("userId") REFERENCES "User" ("id");

ALTER TABLE "Comment" ADD FOREIGN KEY ("documentId") REFERENCES "Document" ("id");

ALTER TABLE "Comment" ADD FOREIGN KEY ("parentCommentId") REFERENCES "Comment" ("id");

ALTER TABLE "Position" ADD FOREIGN KEY ("commentsId") REFERENCES "Comment" ("id");

ALTER TABLE "User_Team" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "User_Team" ADD FOREIGN KEY ("team_id") REFERENCES "Team" ("id");

ALTER TABLE "User_Workspace" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "User_Workspace" ADD FOREIGN KEY ("workspace_id") REFERENCES "Workspace" ("id");

ALTER TABLE "User_Document" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "User_Document" ADD FOREIGN KEY ("document_id") REFERENCES "Document" ("id");

ALTER TABLE "Team_Workspace" ADD FOREIGN KEY ("team_id") REFERENCES "Team" ("id");

ALTER TABLE "Team_Workspace" ADD FOREIGN KEY ("workspace_id") REFERENCES "Workspace" ("id");

ALTER TABLE "Team_Document" ADD FOREIGN KEY ("team_id") REFERENCES "Team" ("id");

ALTER TABLE "Team_Document" ADD FOREIGN KEY ("document_id") REFERENCES "Document" ("id");

ALTER TABLE "Workspace_Document" ADD FOREIGN KEY ("workspace_id") REFERENCES "Workspace" ("id");

ALTER TABLE "Workspace_Document" ADD FOREIGN KEY ("document_id") REFERENCES "Document" ("id");

ALTER TABLE "Document_Collection" ADD FOREIGN KEY ("document_id") REFERENCES "Document" ("id");

ALTER TABLE "Document_Collection" ADD FOREIGN KEY ("collection_id") REFERENCES "Collection" ("id");

ALTER TABLE "Collection_Member" ADD FOREIGN KEY ("collection_id") REFERENCES "Collection" ("id");

ALTER TABLE "Collection_Member" ADD FOREIGN KEY ("member_id") REFERENCES "Member" ("id");

ALTER TABLE "Document_Tag" ADD FOREIGN KEY ("document_id") REFERENCES "Document" ("id");

ALTER TABLE "Document_Tag" ADD FOREIGN KEY ("tag_id") REFERENCES "Tag" ("id");

ALTER TABLE "Collection_Tag" ADD FOREIGN KEY ("folder_id") REFERENCES "Collection" ("id");

ALTER TABLE "Collection_Tag" ADD FOREIGN KEY ("tag_id") REFERENCES "Tag" ("id");

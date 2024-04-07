
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "EnumCommentStatus" AS ENUM (
  'Done',
  'Todo',
  'Doing'
);

CREATE TYPE "EnumActionEventAction" AS ENUM ( 'Download',
	'Upload',
	'Edit',
	'Delete',
	'AddRole',
	'RemoveRole',
	'Duplicate',
	'Comment',
	'Share',
	'Assign',
	'Star',
	'Public',
	'Private'
);

CREATE TABLE "Tag" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" Text UNIQUE,
  "description" Text
);

CREATE TABLE "Team" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
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
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" Text,
  "domain" Text,
  "isPublic" Boolean,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "url" Text
);

CREATE TABLE "Profile" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "userId" UUID UNIQUE,
  "fullName" Text,
  "age" Int,
  "address" Text,
  "avatarUrl" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp
);

CREATE TABLE "User" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
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
  "roles" Json,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp
);

CREATE TABLE "Document" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "title" Text,
  "text" Text,
  "emoji" Text,
  "isPublic" Boolean,
  "isFullWidth" Boolean,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "authorId" UUID,
  "teamId" UUID,
  "workspaceId" UUID,
  "collectionId" UUID
);

CREATE TABLE "Collection" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" Text,
  "description" Text,
  "icon" Text,
  "isSaving" Boolean,
  "canShare" Boolean,
  "childCollectionIds" Text,
  "downloadPermission" Json,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "workspaceId" UUID,
  "parentCollectionId" UUID,
  "ownerUserId" UUID
);

CREATE TABLE "Member" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "userId" UUID,
  "collectionId" UUID,
  "role" Json,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp
);

CREATE TABLE "ActionEvent" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "action" EnumActionEventAction,
  "createdAt" Timestamp DEFAULT (now()),
  "actor" UUID,
  "assignee" UUID,
  "assigner" UUID,
  "documentId" UUID,
  "collectionId" UUID
);

CREATE TABLE "Comment" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "content" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "userId" UUID,
  "documentId" UUID,
  "parentCommentId" UUID
);

CREATE TABLE "Position" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "line" Int,
  "col" Int,
  "toLine" Text,
  "toCol" Text,
  "content" Text,
  "createdAt" Timestamp DEFAULT (now()),
  "updatedAt" Timestamp,
  "commentsId" UUID
);

CREATE TABLE "User_Team" (
  "user_id" UUID,
  "team_id" UUID
);

CREATE TABLE "User_Workspace" (
  "user_id" UUID,
  "workspace_id" UUID
);

CREATE TABLE "User_Document" (
  "user_id" UUID,
  "document_id" UUID
);

CREATE TABLE "Team_Workspace" (
  "team_id" UUID,
  "workspace_id" UUID
);

CREATE TABLE "Team_Document" (
  "team_id" UUID,
  "document_id" UUID
);

CREATE TABLE "Workspace_Document" (
  "workspace_id" UUID,
  "document_id" UUID
);

CREATE TABLE "Document_Collection" (
  "document_id" UUID,
  "collection_id" UUID
);

CREATE TABLE "Collection_Member" (
  "collection_id" UUID,
  "member_id" UUID
);

CREATE TABLE "Document_Tag" (
  "document_id" UUID,
  "tag_id" UUID
);

CREATE TABLE "Collection_Tag" (
  "folder_id" UUID,
  "tag_id" UUID
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

-- Drop foreign key constraints

ALTER TABLE "User_Workspace" DROP CONSTRAINT "FK_User_Workspace_user_id_User_id";
ALTER TABLE "User_Workspace" DROP CONSTRAINT "FK_User_Workspace_workspace_id_Workspace_id";

ALTER TABLE "User_Document" DROP CONSTRAINT "FK_User_Document_user_id_User_id";
ALTER TABLE "User_Document" DROP CONSTRAINT "FK_User_Document_document_id_Document_id";

ALTER TABLE "Member_Team" DROP CONSTRAINT "FK_Member_Team_member_id_Member_id";
ALTER TABLE "Member_Team" DROP CONSTRAINT "FK_Member_Team_team_id_Team_id";

ALTER TABLE "Collection_Member" DROP CONSTRAINT "FK_Collection_Member_collection_id_Collection_id";
ALTER TABLE "Collection_Member" DROP CONSTRAINT "FK_Collection_Member_member_id_Member_id";

ALTER TABLE "Document_Tag" DROP CONSTRAINT "FK_Document_Tag_document_id_Document_id";
ALTER TABLE "Document_Tag" DROP CONSTRAINT "FK_Document_Tag_tag_id_Tag_id";

ALTER TABLE "Collection_Tag" DROP CONSTRAINT "FK_Collection_Tag_folder_id_Collection_id";
ALTER TABLE "Collection_Tag" DROP CONSTRAINT "FK_Collection_Tag_tag_id_Tag_id";

ALTER TABLE "Member" DROP CONSTRAINT "FK_Member_documentId_Document_id";
ALTER TABLE "Member" DROP CONSTRAINT "FK_Member_workspaceId_Workspace_id";

ALTER TABLE "ActionEvent" DROP CONSTRAINT "FK_ActionEvent_documentId_Document_id";
ALTER TABLE "ActionEvent" DROP CONSTRAINT "FK_ActionEvent_collectionId_Collection_id";

ALTER TABLE "Comment" DROP CONSTRAINT "FK_Comment_userId_User_id";
ALTER TABLE "Comment" DROP CONSTRAINT "FK_Comment_documentId_Document_id";
ALTER TABLE "Comment" DROP CONSTRAINT "FK_Comment_parentCommentId_Comment_id";

ALTER TABLE "User" DROP CONSTRAINT "FK_User_id_Member_userId";

ALTER TABLE "Collection" DROP CONSTRAINT "FK_Collection_ownerUserId_User_id";
ALTER TABLE "Collection" DROP CONSTRAINT "FK_Collection_parentCollectionId_Collection_id";
ALTER TABLE "Collection" DROP CONSTRAINT "FK_Collection_workspaceId_Workspace_id";

ALTER TABLE "Document" DROP CONSTRAINT "FK_Document_collectionId_Collection_id";
ALTER TABLE "Document" DROP CONSTRAINT "FK_Document_teamId_Team_id";
ALTER TABLE "Document" DROP CONSTRAINT "FK_Document_workspaceId_Workspace_id";

-- Drop tables

DROP TABLE IF EXISTS "Member";
DROP TABLE IF EXISTS "ActionEvent";
DROP TABLE IF EXISTS "Comment";
DROP TABLE IF EXISTS "Position";
DROP TABLE IF EXISTS "User_Document";
DROP TABLE IF EXISTS "User_Workspace";
DROP TABLE IF EXISTS "Member_Team";
DROP TABLE IF EXISTS "Collection_Member";
DROP TABLE IF EXISTS "Document_Tag";
DROP TABLE IF EXISTS "Collection_Tag";
DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS "Collection";
DROP TABLE IF EXISTS "Document";
DROP TABLE IF EXISTS "Workspace";
DROP TABLE IF EXISTS "Team";
DROP TABLE IF EXISTS "Tag";

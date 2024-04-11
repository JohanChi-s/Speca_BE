package v1

const (
	AdminRole  = 100
	OwnerRole  = 80
	MemberRole = 50
	EditorRole = 20
	ViewerRole = 5
	GuestRole  = 1
)

// RoleToString converts Role to lowercase string
func RoleToString(role int) string {
	switch role {
	case AdminRole:
		return "admin"
	case OwnerRole:
		return "owner"
	case MemberRole:
		return "member"
	case EditorRole:
		return "editor"
	case ViewerRole:
		return "viewer"
	case GuestRole:
		return "guest"
	default:
		return "unknown"
	}
}

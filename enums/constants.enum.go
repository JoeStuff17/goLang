package enums

type UserRole string
type AdminDesignation string
type GenericStatus string

func (u UserRole) String() {
	panic("unimplemented")
}

const (
	RoleSuperAdmin  UserRole = "super_admin"
	RoleAdmin       UserRole = "admin"
	RoleClient      UserRole = "client"
	RoleClientAdmin UserRole = "client_admin"
	RoleClientUser  UserRole = "client_user"
)

const (
	DesignationManager  AdminDesignation = "manager"
	DesignationAdmin    AdminDesignation = "admin"
	DesignationLead     AdminDesignation = "lead"
	DesignationEmployee AdminDesignation = "employee"
)

const (
	NEW      GenericStatus = "0"
	ACTIVE   GenericStatus = "1"
	INACTIVE GenericStatus = "2"
	BLOCKED  GenericStatus = "3"
)

package enums

type UserRole string
type GenericStatus string
type AdminDesignation string
type ChurchDesignation string
type Gender string
type MaritalStatus string
type Professions string

func (u UserRole) String() {
	panic("unimplemented")
}

const (
	RoleSuperAdmin  UserRole = "super_admin"
	RoleAdmin       UserRole = "admin"
	RoleChurch      UserRole = "church"
	RoleChurchAdmin UserRole = "church_admin"
	RoleChurchUser  UserRole = "church_user"
)

const (
	DesignationManager  AdminDesignation = "manager"
	DesignationAdmin    AdminDesignation = "admin"
	DesignationLead     AdminDesignation = "lead"
	DesignationEmployee AdminDesignation = "employee"
)
const (
	ChurchUserRolePastor              ChurchDesignation = "pastor"
	ChurchUserRoleAssistantPastor     ChurchDesignation = "assistant_pastor"
	ChurchUserRolePriest              ChurchDesignation = "priest"
	ChurchUserRoleAdmin               ChurchDesignation = "admin"
	ChurchUserRoleTreasurer           ChurchDesignation = "treasurer"
	ChurchUserRoleSecretary           ChurchDesignation = "secretary"
	ChurchUserRoleElder               ChurchDesignation = "elder"
	ChurchUserRoleDeacon              ChurchDesignation = "deacon"
	ChurchUserRoleYouthLeader         ChurchDesignation = "youth_leader"
	ChurchUserRoleChoirLeader         ChurchDesignation = "choir_leader"
	ChurchUserRoleSundaySchoolTeacher ChurchDesignation = "sunday_school_teacher"
	ChurchUserRoleMember              ChurchDesignation = "member"
)

const (
	NEW      GenericStatus = "0"
	ACTIVE   GenericStatus = "1"
	INACTIVE GenericStatus = "2"
	BLOCKED  GenericStatus = "3"
)

const (
	Male   Gender = "male"
	Female Gender = "female"
	Trans  Gender = "trans"
)

const (
	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
	widowed  MaritalStatus = "widowed"
)

const (
	ProfessionNone               Professions = "none"
	ProfessionStudent            Professions = "student"
	ProfessionUnemployed         Professions = "unemployed"
	ProfessionGovernmentEmployee Professions = "government_employee"
	ProfessionPrivateEmployee    Professions = "private_employee"
	ProfessionBusinessOwner      Professions = "business_owner"
	ProfessionRetired            Professions = "retired"
	ProfessionSelfEmployed       Professions = "self_employed"
	ProfessionHomemaker          Professions = "homemaker"
)

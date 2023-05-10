package services

const (
	SetCommand    = "\n- Set login and password command:\n \t`/set <service> <login> <password>`"
	GetCommand    = "\n- Get login and password command:\n \t`/get <service>`"
	DelCommand    = "\n- Delete login and password command:\n \t`/del <service>`"
	SetSuccessMsg = "Set Success!"
	GetSuccessMsg = "Get Success!"
	DelSuccessMsg = "Delete Success!"
	StartMsg      = "Hello! I am an app for VK Internship. I am save your passwords and logins. Created by @busyleap" +
		"You can use my commands:\n\n1) Setting the login and password for the service.\n" +
		"Example:\n\n`/set <service name> <login> <password>` \n\n2) Obtaining a login and" +
		" password by the name of the service.\nExample:\n\n`/get <service name>` \n\n" +
		"3) Removing the login and password for the service by name:\n\n`/del <service name>` \n\n"
	ParamsCountErrorMsg       = "Incorrect params count!"
	SetDataErrorMsg           = "Failed to set data for the service"
	GetDataErrorMsg           = "Failed to get data for the service"
	DelDataErrorMsg           = "Failed to del data for the service"
	GetPasswordDecodeErrorMsg = "Failed to get password!"
	UnknownCommandErrorMsg    = "Unknown command! Try again with allowed commands:\n" +
		SetCommand + "\n" + GetCommand + "\n" + DelCommand
)

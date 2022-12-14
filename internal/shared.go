package internal

import corev1 "k8s.io/api/core/v1"

// Tenancy labels
const (
	ORGANIZATION_LABEL_KEY      = "robolaunch.io/organization"
	TEAM_LABEL_KEY              = "robolaunch.io/team"
	REGION_LABEL_KEY            = "robolaunch.io/region"
	CLOUD_INSTANCE_LABEL_KEY    = "robolaunch.io/cloud-instance"
	PHYSICAL_INSTANCE_LABEL_KEY = "robolaunch.io/physical-instance"
)

// Ready robot label
const (
	ROBOT_IMAGE_USER       = "robolaunch.io/robot-image-user"
	ROBOT_IMAGE_REPOSITORY = "robolaunch.io/robot-image-repository"
	ROBOT_IMAGE_TAG        = "robolaunch.io/robot-image-tag"
)

// Target robot label
const (
	TARGET_ROBOT = "robolaunch.io/target-robot"
)

// Robot owned resources' postfixes
const (
	PVC_VAR_POSTFIX          = "-pvc-var"
	PVC_ETC_POSTFIX          = "-pvc-etc"
	PVC_OPT_POSTFIX          = "-pvc-opt"
	PVC_USR_POSTFIX          = "-pvc-usr"
	PVC_DISPLAY_POSTFIX      = "-pvc-display"
	PVC_WORKSPACE_POSTFIX    = "-pvc-workspace"
	DISCOVERY_SERVER_POSTFIX = "-discovery"
	JOB_LOADER_POSTFIX       = "-loader"
	ROS_BRIDGE_POSTFIX       = "-bridge"
)

// BuildManager owned resources' postfixes
const (
	CONFIG_MAP_SCRIPTS = "-scripts"
)

// LaunchManager owned resources' postfixes
const (
	POD_LAUNCH_POSTFIX = "-launch"
)

// RobotVDI owned resources' postfixes
const (
	PVC_VDI_POSTFIX     = "-display"
	SVC_TCP_VDI_POSTFIX = "-tcp"
	SVC_UDP_VDI_POSTFIX = "-udp"
	POD_VDI_POSTFIX     = ""
	INGRESS_VDI_POSTFIX = ""
)

// Paths

const (
	CUSTOM_SCRIPTS_PATH = "/etc/custom"
	HELPERS_PATH        = "/var/lib/robolaunch-helpers/"
	X11_UNIX_PATH       = "/tmp/.X11-unix"
)

// Super client configuration
const (
	SUPER_CLIENT_CONFIG = "<?xml version='1.0' encoding='UTF-8' ?>" +
		"<dds>" +
		"<profiles xmlns='http://www.eprosima.com/XMLSchemas/fastRTPS_Profiles'>" +
		"<participant profile_name='super_client_profile' is_default_profile='true'>" +
		"<rtps>" +
		"		<builtin>" +
		"			<discovery_config>" +
		"				<discoveryProtocol>SUPER_CLIENT</discoveryProtocol>" +
		"				<discoveryServersList>" +
		"					<RemoteServer prefix='44.53.00.5f.45.50.52.4f.53.49.4d.41'>" +
		"						<metatrafficUnicastLocatorList>" +
		"							<locator>" +
		"								<udpv4>" +
		"								<address>" + "%s" + "</address>" +
		"									<port>11811</port>" +
		"								</udpv4>" +
		"							</locator>" +
		"						</metatrafficUnicastLocatorList>" +
		"					</RemoteServer>" +
		"				</discoveryServersList>" +
		"			</discovery_config>" +
		"		</builtin>" +
		"	</rtps>" +
		"</participant>" +
		"</profiles>" +
		"</dds>"
)

// Ingress annotations
const (
	INGRESS_VDI_AUTH_URL_KEY              = "nginx.ingress.kubernetes.io/auth-url"
	INGRESS_VDI_AUTH_URL_VAL              = "%s"
	INGRESS_VDI_AUTH_SIGNIN_KEY           = "nginx.ingress.kubernetes.io/auth-signin"
	INGRESS_VDI_AUTH_SIGNIN_VAL           = "%s"
	INGRESS_VDI_AUTH_RESPONSE_HEADERS_KEY = "nginx.ingress.kubernetes.io/auth-response-headers"
	INGRESS_VDI_AUTH_RESPONSE_HEADERS_VAL = "x-auth-request-user, x-auth-request-email, x-auth-request-access-token"
	INGRESS_VDI_CONFIGURATION_SNIPPET_KEY = "nginx.ingress.kubernetes.io/configuration-snippet"
	INGRESS_VDI_CONFIGURATION_SNIPPET_VAL = "" +
		"        #proxy_set_header Host $host;\n" +
		"		proxy_set_header X-Real-IP $remote_addr;\n" +
		"		proxy_set_header X-Forwarded-For $remote_addr;\n" +
		"		proxy_set_header X-Forwarded-Host $host;\n" +
		"		proxy_set_header X-Forwarded-Port $server_port;\n" +
		"		proxy_set_header X-Forwarded-Protocol $scheme;\n"
	INGRESS_VDI_CERT_MANAGER_KEY               = "acme.cert-manager.io/http01-edit-in-place"
	INGRESS_VDI_CERT_MANAGER_VAL               = "true"
	INGRESS_VDI_NGINX_PROXY_BUFFER_SIZE_KEY    = "nginx.ingress.kubernetes.io/proxy-buffer-size"
	INGRESS_VDI_NGINX_PROXY_BUFFER_SIZE_VAL    = "16k"
	INGRESS_VDI_NGINX_PROXY_BUFFERS_NUMBER_KEY = "nginx.ingress.kubernetes.io/proxy-buffers-number"
	INGRESS_VDI_NGINX_PROXY_BUFFERS_NUMBER_VAL = "4"
	INGRESS_VDI_NGINX_REWRITE_TARGET_KEY       = "nginx.ingress.kubernetes.io/rewrite-target"
	INGRESS_VDI_NGINX_REWRITE_TARGET_VAL       = "/$2"
)

func Bash(command string) []string {
	return []string{
		"/bin/bash",
		"-c",
		command,
	}
}

func Env(key string, value string) corev1.EnvVar {
	return corev1.EnvVar{
		Name:  key,
		Value: value,
	}
}

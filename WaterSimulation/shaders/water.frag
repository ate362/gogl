#version 410 core
#define SPECULAR_STRENGTH 2

const vec3 WATER_BLUE = vec3(64.0f/255.0f, 164.0f/255.0f, 223.0f/255.0f);

const vec3 ambient = vec3(0.28f, 0.3f, 0.25f);
const vec3 diffuse = WATER_BLUE;//vec3(0.25, 0.0, 0.0);
const vec3 specular = vec3(1.0, 1.0, 1.0);
const float shininess = 20.0;
const vec4 lightColor = vec4(1.0, 1.0, 1.0, 1.0);

uniform vec3 lightPos;
uniform vec3 eyePosition;

out vec4 color;

vec4 iamb, idiff, ispec;	

in VS_OUT
{
	vec4 position;
	vec4 normal;
	vec2 texCoords;
} fs_in;


void Phong(inout vec4 iamb, inout vec4 idiff, inout vec4 ispec)
{
    vec4 normal = normalize(fs_in.normal);
	
	vec4 lightVector, viewVector, reflectVector;
	float diff, spec;

	lightVector = normalize(vec4(lightPos, 1.0f) - fs_in.position);
	diff = max(dot(normal, lightVector), 0.0f);

	viewVector    = normalize(vec4(eyePosition, 1.0f) - fs_in.position);
	reflectVector = normalize(reflect(-lightVector, normal));
	spec = pow(max(dot(viewVector, reflectVector), 0.0f), shininess);

	iamb   += 	   vec4(ambient, 1.0f);
	idiff  += diff * vec4(diffuse, 1.0f) * lightColor;
	ispec  += SPECULAR_STRENGTH * spec * vec4(specular, 1.0f);	
}

void main()
{
	iamb = idiff = ispec = vec4(0.0f);

	Phong(iamb, idiff, ispec);

    // the color of the light "reflects" off the object
    color = (iamb+idiff)*vec4(WATER_BLUE, 1f) + ispec;
}
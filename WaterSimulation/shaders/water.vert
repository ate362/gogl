#version 410 core
#extension GL_ARB_explicit_uniform_location : enable
#define PI 3.14159f
#define g 9.81f
#define SIZE 6

// Input attributes base: 0
layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normal;
layout (location = 2) in vec2 texCoords;
layout (location = 3) in vec3 tangent;

uniform mat4 world;
uniform mat4 camera;
uniform mat4 project;
uniform mat4 inverseTranspose;

struct WaterDeformer
{
    float time;
    int numWaves;
    float amplitude[SIZE];
    float wavelength[SIZE];
    float speed[SIZE];
    vec2 direction[SIZE];
};
uniform WaterDeformer waterDeformer;

float wave(int i, float x, float z);
float waveHeight(float x, float z);
float dWavedx(int i, float x, float z);
float dWavedz(int i, float x, float z);
vec3 waveNormal(float x, float z);

out VS_OUT
{
	vec4 position;
	vec4 normal;
	vec2 texCoords;
} vs_out;

float rand(vec2 co){
    return fract(sin(dot(co, vec2(12.9898, 78.233))) * 43758.5453) * 0.5 - 1.0;
}

void main()
{	
	vec3 deformedPosition = position;
    vec3 deformedNormal = waveNormal(deformedPosition.x, deformedPosition.z);

    deformedPosition.y = waveHeight(deformedPosition.x, deformedPosition.z); 

    vs_out.position  = world * vec4(deformedPosition, 1.0f);
	vs_out.normal 	 = inverseTranspose * vec4(deformedNormal, 0.0f);
	vs_out.texCoords = texCoords;

    gl_Position = project * camera * world * vec4(deformedPosition, 1.0f);        
}


float wave(int i, float x, float z)
{
    float frequency = 2 * PI / waterDeformer.wavelength[i];
    float phase = waterDeformer.speed[i] * frequency;
    //float theta = dot(waterDeformer.direction[i], vec2(x, z));
    float theta = -dot(normalize(vec2(x, z) - vec2(0, 0)),vec2(x,z));
    return waterDeformer.amplitude[i] * sin(theta * frequency + waterDeformer.time * phase);
}

float waveHeight(float x, float z)
{
    float height = 0.0f;
    for (int i = 0; i < waterDeformer.numWaves; ++i)
        height += wave(i, x, z);   
    float h = rand(vec2(x, z));     
    return height;//+h*.07f;
}

float dWavedx(int i, float x, float z)
{
    float frequency = 2 * PI / waterDeformer.wavelength[i];
    float phase = waterDeformer.speed[i] * frequency;
    //float theta = dot(waterDeformer.direction[i], vec2(x, z));
    float theta = -dot(normalize(vec2(x, z) - vec2(0, 0)),vec2(x,z));
    float A = waterDeformer.amplitude[i] * waterDeformer.direction[i].x * frequency;
    return A * cos(theta * frequency + waterDeformer.time * phase);
}

float dWavedz(int i, float x, float z)
{
    float frequency = 2 * PI / waterDeformer.wavelength[i];
    float phase = waterDeformer.speed[i] * frequency;
    //float theta = dot(waterDeformer.direction[i], vec2(x, z));
    float theta = -dot(normalize(vec2(x, z) - vec2(0, 0)),vec2(x,z));
    float A = waterDeformer.amplitude[i] * waterDeformer.direction[i].y * frequency;
    return A * cos(theta * frequency + waterDeformer.time * phase);
}

vec3 waveNormal(float x, float z)
{
    float dx = 0.0f;
    float dz = 0.0f;
    for (int i = 0; i < waterDeformer.numWaves; ++i)
    {
        dx += dWavedx(i, x, z);
        dz += dWavedz(i, x, z);
    }
    float h = rand(vec2(x, z));
    vec3 n = vec3(dx, 1, dz);
    return normalize(n);
}
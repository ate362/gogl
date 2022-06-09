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

out vec2 TexCoord;

void main()
{	
	vec3 deformedPosition = position;

    deformedPosition.y = waveHeight(deformedPosition.x, deformedPosition.z);    
    gl_Position = project * camera * world * vec4(deformedPosition, 1.0f);        
}


float wave(int i, float x, float z)
{
    float frequency =  2 * PI / waterDeformer.wavelength[i];
    float phase = waterDeformer.speed[i] * frequency;
    // float theta = dot(waterDeformer.direction[i], vec2(x, z));
    float theta = -dot(normalize(vec2(x, z) - vec2(0, 0)),vec2(x,z));
    return waterDeformer.amplitude[i] * sin(theta * frequency + waterDeformer.time * phase);
}

float waveHeight(float x, float z)
{
    float height = 0.0f;
    for (int i = 0; i < waterDeformer.numWaves; ++i)
        height += wave(i, x, z);        
    return height;
}
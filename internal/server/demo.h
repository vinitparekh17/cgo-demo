#include <stdio.h>
#include <stdlib.h>

const char* greet(char* name) {
    // create an array to randomize the greetings message
    const char* greetings[] = {
        "Hello",
        "Hi",
        "Hey",
        "Greetings",
        "Good day",
        "Good morning",
        "Good afternoon",
        "Good evening",
        "Good night",
        "Howdy",
        "Salutations",
        "What's up",
        "Yo",
        "Hiya",
        "Hey there",
        "Hey you",
        "Hey buddy",
        "Hey friend",
        "Hey mate",
        "Hey pal",
        "Hey dude"
    };

    // get the size of the greetings array
    int greetings_size = sizeof(greetings) / sizeof(greetings[0]);

    // get a random index
    int random_index = rand() % greetings_size;

    char* message = (char*)malloc(100 * sizeof(char));
    sprintf(message, "%s, %s!", greetings[random_index], name);

    return message;
}
#ifndef BUTTON_H_
#define BUTTON_H_

#include "../include/raylib.h"

typedef struct
{
  int width;
  int height;
  Color color;
  char *body;
} Button;

#endif // BUTTON_H_
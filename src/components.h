#ifndef COMPONENTS_H
#define COMPONENTS_H

#include "raylib.h"

struct Div;
struct P;

typedef void (*Div_click)(struct Div *);
typedef void (*P_click)(struct P *);

#define DEFINE_ON(Type)    \
  typedef struct On_##Type \
  {                        \
    Type##_click click;    \
  } On_##Type

DEFINE_ON(Div);
DEFINE_ON(P);

typedef struct P
{
  Color color;
  Color *background_color;
  int font_size;
  char *text;
  On_P *On_P;
} P;

typedef struct
{
  int left;
  int right;
  int top;
  int bottom;
} Margin;

typedef struct Div
{
  int width;
  int height;
  Color color;
  char *body;
  On_Div on;
  Margin margin;
} Div;

// Function prototype
void print_div(const Div *div);

#endif // COMPONENTS_H

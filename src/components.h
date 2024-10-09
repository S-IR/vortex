#ifndef COMPONENTS_H
#define COMPONENTS_H

#include "raylib.h"
#include <stddef.h>

struct Div;
struct P;
struct Component;

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
  int x;
  int y;
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
  int x;
  int y;
  int width;
  int height;
  Color color;
  struct Component *body;
  size_t body_len;
  Margin margin;
  On_Div on;
} Div;

// Function prototype
void print_div(const Div *div);

typedef enum
{
  TYPE_DIV,
  TYPE_P,
} COMPONENT_TYPE;

typedef union
{
  Div div;
  P p;
} ComponentBody;

typedef struct
{
  COMPONENT_TYPE type;
  ComponentBody body;
} Component;

#endif // COMPONENTS_H

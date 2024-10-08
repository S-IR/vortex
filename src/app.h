#ifndef APP_H_
#define APP_H_

#include <stddef.h>
#include "components.h"

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

#define MAX_COMPONENTS 1024

#define App_add(app, x) _Generic((x), \
    Div: App_add_div,                 \
    P: App_add_p)(app, x)

typedef struct
{
  int width;
  int height;
  char *title;
  int currX;
  int currY;
  Component body[MAX_COMPONENTS];
  size_t body_size;
} App;

void App_start(App *app);
void App_add_div(App *app, Div newDiv);
void App_add_p(App *app, P new_p);

#endif // APP_H_

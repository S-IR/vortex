#include "components.h"
#include "raylib.h" // Make sure to include Raylib header
#include <stdio.h>
#include <assert.h>
#include <stddef.h>
#include "components.h"
#include "app.h"

#define LEN(x) ((sizeof(x) / sizeof(0 [x])) / ((size_t)(!(sizeof(x) % sizeof(0 [x])))))

static void draw_div(Div *div, App *app);
static void draw_p(P *p, App *app);

#define draw(x, app) _Generic((x), \
    Div *: draw_div,               \
    P *: draw_p)(x, app)

void App_start(App *app)
{
  // Initialize the Raylib window
  InitWindow(app->width, app->height, app->title);
  SetTargetFPS(60); // Set the target FPS

  while (!WindowShouldClose())
  {
    // Draw
    BeginDrawing();
    ClearBackground(RAYWHITE);

    for (size_t i = 0; i < LEN(app->body); i++)
    {
      Component comp = app->body[i];

      switch (comp.type)
      {
      case TYPE_DIV:
        draw(&comp.body.div, app);
        break;
      case TYPE_P:
        draw(&comp.body.p, app);
        break;
      }
    }
    app->currX = 0;
    app->currY = 0;
    EndDrawing();
  }
}
void App_add_div(App *app, Div newDiv)
{
  assert(app->body_size < MAX_COMPONENTS && "MAX APP SIZE REACHED");

  app->body[app->body_size]
      .type = TYPE_DIV;
  app->body[app->body_size].body.div = newDiv;
  app->body_size++;
}

void App_add_p(App *app, P new_p)
{
  assert(app->body_size < MAX_COMPONENTS && "MAX APP SIZE REACHED");

  app->body[app->body_size]
      .type = TYPE_P;
  app->body[app->body_size].body.p = new_p;
  app->body_size++;
}

static void draw_div(Div *div, App *app)
{
  assert(div != NULL && app != NULL);

  DrawRectangle(app->currX, app->currY, div->width, div->height, div->color);
  app->currX = 0;
  app->currY += div->height;
}

static void draw_p(P *p, App *app)
{
  assert(p != NULL && app != NULL);
  DrawText(p->text, app->currX, app->currY, p->font_size, p->color);
  int textWidth = MeasureText(p->text, p->font_size);
  // int textHeight = p->font_size;
  app->currX += textWidth;
  // app->currY += textHeight;
}

// Call this to close the window when done
void App_close()
{
  CloseWindow();
}

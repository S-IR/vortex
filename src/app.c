#include "components.h"
#include "raylib.h" // Make sure to include Raylib header
#include <stdio.h>
#include <assert.h>
#include <stddef.h>
#include "components.h"
#include "app.h"
#include <stddef.h>
#include "debug.h"

#define LEN(x) ((sizeof(x) / sizeof(0 [x])) / ((size_t)(!(sizeof(x) % sizeof(0 [x])))))

static void handle_click(App *app);

static void draw_div(Div *div, App *app);
static void draw_p(P *p, App *app);

#define draw(x, app) _Generic((x), \
    Div *: draw_div,               \
    P *: draw_p)(x, app)

void App_start(App *app)
{

  SetTargetFPS(60); // Set the target FPS
  SetConfigFlags(FLAG_WINDOW_RESIZABLE);
  // Initialize the Raylib window
  InitWindow(app->width, app->height, app->title);

  while (!WindowShouldClose())
  {
    // Draw
    BeginDrawing();

    ClearBackground(RAYWHITE);
    app->currX = 0;
    app->currY = 0;

    if (IsKeyPressed(KEY_F9))
    {
    }

    for (size_t i = 0; i < app->body_size; i++)
    {
      Component *comp = &app->body[i];

      switch (comp->type)
      {
      case TYPE_DIV:
        draw(&comp->body.div, app);
        break;
      case TYPE_P:
        draw(&comp->body.p, app);
        break;
      }
    }
    if (IsMouseButtonPressed(MOUSE_BUTTON_LEFT))
    {
      handle_click(app);
    }

    setup_debug();

    EndDrawing();
  }
}
void App_add_div(App *app, Div newDiv)
{
  assert(app->body_size < MAX_COMPONENTS && "MAX APP SIZE REACHED");
  assert(newDiv.width >= 0 && "ERROR: width of the new div is too small");
  assert(newDiv.height >= 0 && "ERROR: height of the new div is too small");

  app->body[app->body_size]
      .type = TYPE_DIV;
  app->body[app->body_size].body.div = newDiv;
  // printf("Adding div: width %d, height %d\n", newDiv.width, newDiv.height);

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
  assert(div->width >= 0 && "DIV WIDTH IS 0");

  div->x = app->currX;
  div->y = app->currY;
  DrawRectangle(div->x, div->y, div->width, div->height, div->color);

  // printf("from drawing : x %x y %d width %d height %d\n", div->x, div->y, div->width, div->height);
  app->currX = 0;
  app->currY += div->height;
}

static void draw_p(P *p, App *app)
{
  assert(p != NULL && app != NULL);
  DrawText(p->text, app->currX, app->currY, p->font_size, p->color);
  int textWidth = MeasureText(p->text, p->font_size);
  // int textHeight = p->font_size;
  p->x = app->currX;
  p->y = app->currY;

  app->currX += textWidth;
  // app->currY += textHeight;
}
// In the handle_click function:
static void handle_click(App *app)
{
  printf("Started click\n");

  Vector2 mousePos = GetMousePosition();
  for (size_t i = 0; i < app->body_size; i++)
  {
    Component *comp = &app->body[i]; // Use pointer to avoid copying
    switch (comp->type)
    {
    case TYPE_DIV:
    {
      Div *curr_div = &comp->body.div;
      printf("Checking div: x %d, y %d, width %d, height %d\n", curr_div->x, curr_div->y, curr_div->width, curr_div->height);
      Rectangle area = {
          .x = curr_div->x,
          .y = curr_div->y,
          .width = curr_div->width,
          .height = curr_div->height};

      if (CheckCollisionPointRec(mousePos, area))
      {
        printf("Click detected on div: x %d, y %d, width %d, height %d\n", curr_div->x, curr_div->y, curr_div->width, curr_div->height);
        if (is_debug_mode())
        {
          debug_select_component(comp);
        }

        if (curr_div->on.click != NULL)
        {
          curr_div->on.click(curr_div);
        }
      }
      break;
    }
    default:
      break;
    }
  }
}

// Call this to close the window when done
void App_close()
{
  CloseWindow();
}

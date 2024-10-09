#include "debug.h"
#include <stdbool.h>
#include <stdio.h>
#include "components.h"
#include <string.h>

static Component *selectedComponent = NULL;

static bool debug_mode = false;

// double last_time = 0.0;
// const double delay = 0.1;

void setup_debug()
{
  // double curr_time = GetTime();

  if (IsKeyDown(KEY_F9))
  {
    printf("pressing f9\n");
  }
  // Toggle the debug mode if F9 is pressed and the delay has passed
  if (IsKeyPressed(KEY_F9))
  {
    debug_mode = !debug_mode;
    // last_time = curr_time;
  }

  if (!debug_mode)
    return;
  int screenWidth = GetScreenWidth();
  int screenHeight = GetScreenHeight();
  int sidebarWidth = 200;

  int sidebarStart = screenWidth - sidebarWidth;
  DrawRectangle(screenWidth - sidebarWidth, 0, sidebarWidth, screenHeight, GRAY);

  int textWidth = MeasureText("Debug Mode", 20);
  int finalY = sidebarStart + (sidebarWidth / 2) - (textWidth / 2);

  int currY = 10;
  DrawText("Debug Mode", finalY, currY, 20, RED);
  currY += 10;
  if (selectedComponent != NULL)
  {
    switch (selectedComponent->type)
    {
    case TYPE_DIV:
    {
      Div curr_div = selectedComponent->body.div;
      DrawText("Div:", sidebarStart + 10, currY, 20, BLUE);

      currY += 10;
      char *widthLabel = "Width :";
      int size = snprintf(NULL, 0, "%s%d", widthLabel, curr_div.width) + 1;
      char widthText[size];
      snprintf(widthText, size, "%s%d", widthLabel, curr_div.width);
      DrawText(widthText, sidebarStart + 10, currY, 20, BLUE);

      currY += 10;
      char *heightLabel = "Height :";
      size = snprintf(NULL, 0, "%s%d", heightLabel, curr_div.height) + 1;
      char heightText[size];
      snprintf(heightText, size, "%s%d", heightLabel, curr_div.height);
      DrawText(heightText, sidebarStart + 10, currY, 20, BLUE);

      currY += 10;
      char *xLabel = "X :";
      size = snprintf(NULL, 0, "%s%d", xLabel, curr_div.x) + 1;
      char xText[size];
      snprintf(xLabel, size, "%s%d", xLabel, curr_div.x);
      DrawText(xText, sidebarStart + 10, currY, 20, BLUE);

      currY += 10;
      char *yLabel = "X :";
      size = snprintf(NULL, 0, "%s%d", yLabel, curr_div.y) + 1;
      char yText[size];
      snprintf(yLabel, size, "%s%d", yLabel, curr_div.y);
      DrawText(yText, sidebarStart + 10, currY, 20, BLUE);
    }
    break;

    default:
      break;
    }
  }
}

bool is_debug_mode()
{
  return debug_mode;
}
void debug_select_component(Component *comp)
{
  selectedComponent = comp;
}
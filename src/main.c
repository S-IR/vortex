#include "../include/raylib.h"
#include <stdlib.h>
#include <time.h>
#include "app.h"
#include <stdio.h>
#include "app.h"
#include "components.h"

int currX = 0;
int currY = 0;

void handle_div_click(Div *div)
{
  printf("Hello world\n");
}

int main()
{
  srand(time(NULL));
  App app = {
      .width = 800,
      .height = 600,
      .title = "hello world",
  };

  for (size_t i = 0; i < 5; i++)
  {

    On_Div callbacks = {.click = handle_div_click};
    Div newDiv = {
        .width = (rand() % 200) + 50,  // Width between 50 and 250
        .height = (rand() % 100) + 25, // Height between 25 and 125
        .body = "hello!",
        .color = {
            .r = rand() % 256, // Red component between 0 and 255
            .g = rand() % 256, // Green component between 0 and 255
            .b = rand() % 256, // Blue component between 0 and 255
            .a = 255           // Fully opaque
        }};

    App_add(&app, newDiv);
  }

  for (size_t i = 0; i < 5; i++)
  {

    Color *bgColor;
    P newP = {
        .background_color = NULL,
        // .color = {
        //     .r = rand() % 256, // Red component between 0 and 255
        //     .g = rand() % 256, // Green component between 0 and 255
        //     .b = rand() % 256, // Blue component between 0 and 255
        //     .a = 255           // Fully opaque
        // },
        .color = BLACK,
        .font_size = (rand() % 50) + 12,
        .text = "my text"};

    App_add(&app, newP);
  }

  App_start(&app);

  return 0;
}

#include <stdio.h>
#include "components.h"

void print_div(const Div *div)
{
  if (div == NULL)
  {
    printf("DIV IS NULL");
    return;
  }

  printf("Div Properties:\n");
  printf("  Width: %d\n", div->width);
  printf("  Height: %d\n", div->height);
  printf("  Color: (R: %d, G: %d, B: %d, A: %d)\n",
         div->color.r, div->color.g, div->color.b, div->color.a);
  // printf("  Body: %s\n", div->body ? div->body : "NULL");
}

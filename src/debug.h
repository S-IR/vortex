#ifndef DEBUG_H
#define DEBUG_H

#include "raylib.h"
#include "components.h"

void setup_debug();

bool is_debug_mode();
void debug_select_component(Component *comp);
#endif // DEBUG_H

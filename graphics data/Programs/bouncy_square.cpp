#include "AllegroUtil.hpp"
#include "screen.hpp"
#include "shapes.hpp"

#include <vector>

const ALLEGRO_COLOR RED   = al_map_rgb(255, 0, 0);
const ALLEGRO_COLOR WHITE = al_map_rgb(255, 255, 255);

screen::Window window(180, 1200, 720);

void NextFrame()
{
    window.ManageCollisions();
    window.MoveAll();
}

void Draw()
{
    al_clear_to_color(WHITE);
    window.DrawAll();
}

int main()
{
    shape::Square sqr1(100, 500, 40, 135);
    shape::Square sqr2(200, 200, 75, 330);

    window.AddObject(&sqr1, RED);
    window.AddObject(&sqr2);
    window.Run(&NextFrame, &Draw);

    return 0;
}
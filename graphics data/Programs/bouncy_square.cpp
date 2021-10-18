#include "screen.hpp"
#include "shapes.hpp"

#include <vector>

#define BLACK     al_map_rgb(0, 0, 0)
#define RED       al_map_rgb(255, 0, 0)
#define GREEN     al_map_rgb(0, 255, 0)
#define BLUE      al_map_rgb(0, 0, 255)
#define DARKRED   al_map_rgb(128, 0, 0)
#define DARKGREEN al_map_rgb(0, 128, 0)
#define DARKBLUE  al_map_rgb(0, 0, 128)
#define PINK      al_map_rgb(255, 20, 147)
#define ORANGE    al_map_rgb(255, 140, 0)
#define YELLOW    al_map_rgb(255, 255, 0)
#define PURPLE    al_map_rgb(147, 112, 219)
#define BROWN     al_map_rgb(160, 82, 45)
#define BEIGE     al_map_rgb(210, 180, 140)
#define LIGHTGRAY al_map_rgb(211, 211, 211)
#define DARKGRAY  al_map_rgb(105, 105, 105)
#define WHITE     al_map_rgb(250, 250, 250)

screen::Window window(210, 1200, 720);

void NextFrame()
{
    window.ManageCollisions();
    window.MoveAll();
}

void Draw()
{
    al_clear_to_color(LIGHTGRAY);
    window.DrawAll();
}

int main()
{
    shape::Square sqr1(100, 220, 40, 135);
    shape::Square sqr2(200, 520, 75, 330);
    shape::Square sqr3(1200, 520, 40, 30); // 1200, 520, 35, 30
    shape::Square sqr4(600, 400, 50, 180);
    shape::Square sqr5(100, 20, 5, 350);
    // shape::Square sqr1(700, 170, 40, 30); // failing setup
    // shape::Square sqr2(800, 120, 75, 60);

    window.AddObject(sqr1, DARKRED); // change init
    window.AddObject(sqr2, ORANGE);
    window.AddObject(sqr3, GREEN);
    window.AddObject(sqr4, BLUE);
    window.AddObject(sqr5, DARKGREEN);
    window.Run(&NextFrame, &Draw);

    return 0;
}
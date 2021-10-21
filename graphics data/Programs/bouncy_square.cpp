#include "screen.hpp"
#include "shapes.hpp"

#include <iostream>
#include <vector>

screen::Window window(210, 1200, 720);

void NextFrame()
{
    window.ManageCollisions();
    window.MoveAll();
}

void Draw()
{
    //al_clear_to_color(LIGHTGRAY);
    window.DrawAll();
}

int main()
{
    shape::Square sqr1(shape::Movement::DYNAMIC, 100, 220, 40, 135);
    shape::Square sqr2(shape::Movement::DYNAMIC, 200, 520, 75, 330);
    shape::Square sqr3(shape::Movement::DYNAMIC, 1200, 520, 40, 30); // 1200, 520, 35, 30
    shape::Square sqr4(shape::Movement::DYNAMIC, 600, 400, 50, 180);
    shape::Square sqr5(shape::Movement::DYNAMIC, 100, 20, 5, 350);
    // shape::Square sqr1(shape::Shape::Movement::DYNAMIC, 700, 170, 40, 30); // failing setup
    // shape::Square sqr2(shape::Shape::Movement::DYNAMIC, 800, 120, 75, 60);

    window.AddObject(sqr1, DARKRED); // change init
    window.AddObject(sqr2, ORANGE);
    window.AddObject(sqr3, GREEN);
    window.AddObject(sqr4, BLUE);
    window.AddObject(sqr5, DARKGREEN);
    window.Run(&NextFrame, &Draw);

    return 0;
}
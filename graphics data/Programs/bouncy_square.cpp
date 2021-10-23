#include "screen.hpp"
#include "shapes.hpp"

#include <iostream>
#include <vector>

screen::Window window(1, 1200, 720);

void NextFrame()
{
    window.ManageCollisions();
    window.MoveAll();
}

void Draw()
{
    window.DrawAll();
}

int main()
{
    // shape::Square sqr1(shape::Movement::DYNAMIC, 100, 220, 40, 135);
    // shape::Square sqr2(shape::Movement::DYNAMIC, 200, 520, 75, 330);
    // shape::Square sqr3(shape::Movement::DYNAMIC, 1100, 520, 40, 30); // 1200, 520, 35, 30
    // shape::Square sqr4(shape::Movement::DYNAMIC, 600, 400, 50, 180);
    // shape::Square sqr5(shape::Movement::DYNAMIC, 100, 20, 5, 350);

    // shape::Square sqr1(shape::Movement::DYNAMIC, 700, 170, 40, 30); // failing setup
    // shape::Square sqr2(shape::Movement::DYNAMIC, 800, 120, 75, 60);

    shape::Square sqr1(shape::Movement::DYNAMIC, 915.294123, 205.150758, 40, 315);
    shape::Square sqr5(shape::Movement::DYNAMIC, 937.101782, 221.99473, 5, 169);
    window.AddObject(sqr1, DARKRED); //pass object initialization
    // window.AddObject(sqr2, ORANGE);
    // window.AddObject(sqr3, GREEN);
    // window.AddObject(sqr4, BLUE);
    window.AddObject(sqr5, DARKGREEN);
    window.Run(&NextFrame, &Draw);

    return 0;
}
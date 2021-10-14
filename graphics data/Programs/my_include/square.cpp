#include "shapes.hpp"

#include <cstdio>

namespace shape
{
    double Square::LeftmostX() const
    {
        return centre.x - edge;
    }

    double Square::RightmostX() const
    {
        return centre.x + edge;
    }

    double Square::UppermostY() const
    {
        return centre.y - edge;
    }

    double Square::LowermostY() const
    {
        return centre.y + edge;
    }

    void Square::SetSides()
    {

        double lx = LeftmostX();
        double rx = RightmostX();
        double uy = UppermostY();
        double ly = LowermostY();

        // printf("side[0]: %.0f %.0f %.0f %.0f\n", lx, uy, lx + edge * 2, uy + 0);
        // printf("side[1]: %.0f %.0f %.0f %.0f\n", rx, uy, rx + 0, uy + edge * 2);
        // printf("side[2]: %.0f %.0f %.0f %.0f\n", lx, uy, lx + 0, uy + edge * 2);
        // printf("side[3]: %.0f %.0f %.0f %.0f\n\n", lx, ly, lx + edge * 2, ly + 0);

        sides[0].SetCoordinates(lx, uy, edge * 2, 0); // top side
        sides[1].SetCoordinates(rx, uy, 0, edge * 2); // right side
        sides[2].SetCoordinates(lx, uy, 0, edge * 2); // left side
        sides[3].SetCoordinates(lx, ly, edge * 2, 0); // bottom side

        printf("POINTER: %p\n", this);
        printf("side[0]: %.0f %.0f %.0f %.0f\n", sides[0].a.x, sides[0].a.y, sides[0].a.x + sides[0].b.x, sides[0].a.y + sides[0].b.y);
        printf("side[1]: %.0f %.0f %.0f %.0f\n", sides[1].a.x, sides[1].a.y, sides[1].a.x + sides[1].b.x, sides[1].a.y + sides[1].b.y);
        printf("side[2]: %.0f %.0f %.0f %.0f\n", sides[2].a.x, sides[2].a.y, sides[2].a.x + sides[2].b.x, sides[2].a.y + sides[2].b.y);
        printf("side[3]: %.0f %.0f %.0f %.0f\n\n", sides[3].a.x, sides[3].a.y, sides[3].a.x + sides[3].b.x, sides[3].a.y + sides[3].b.y);
    }

    Square::Square(Point<double> centreCoords, double side, double alpha) :
        Shape(centreCoords, alpha),
        edge((side < 0) ? 0 : side / 2)
    {
        SetDirection(alpha);
        SetSides();
    }

    Square::Square(double centreX, double centreY, double side, double alpha) :
        Square(Point<double>(centreX, centreY), side, alpha) { }

    const int Square::sideAmount() const
    {
        return 4;
    }

    Shape* Square::InitFromStdin() const
    {
        double x, y;
        double side;
        double angle;

        printf("Enter x-coordinate of square's centre: ");
        scanf("%f\n", x);

        printf("Enter y-coordinate of square's centre: ");
        scanf("%f\n", y);

        printf("Enter square's side: ");
        scanf("%f\n", side);

        printf("Enter square's movement angle: ");
        scanf("%f\n", angle);

        Square sqr = Square(Point<double>(x, y), side, angle);
        return &sqr;
    }

    void Square::Move()
    {
        centre.x += direction.first;
        centre.y -= direction.second;
        printf("POINTER: %p\n", this);
        SetSides();
    }

    void Square::Draw() const
    {
        double lx = this->LeftmostX();
        double rx = this->RightmostX();
        double uy = this->UppermostY();
        double ly = this->LowermostY();
        al_draw_filled_rectangle(lx, uy, rx, ly, color);
    }
}
#include "shapes.hpp"

#include <cstdio>

namespace shape
{
    double Square::LeftmostX() const
    {
        return this->centre.x - edge;
    }

    double Square::RightmostX() const
    {
        return this->centre.x + edge;
    }

    double Square::UppermostY() const
    {
        return this->centre.y - edge;
    }

    double Square::LowermostY() const
    {
        return this->centre.y + edge;
    }

    void Square::SetSides()
    {
        double lx = this->LeftmostX();
        double rx = this->RightmostX();
        double uy = this->UppermostY();
        double ly = this->LowermostY();

        this->sides[0].SetCoordinates(lx, uy, this->edge * 2, 0); //top side
        this->sides[1].SetCoordinates(rx, uy, 0, this->edge * 2); //right side
        this->sides[2].SetCoordinates(lx, uy, 0, this->edge * 2); //left side
        this->sides[3].SetCoordinates(lx, ly, this->edge * 2, 0); //bottom side
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
        this->centre.x += direction.first;
        this->centre.y -= direction.second;
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
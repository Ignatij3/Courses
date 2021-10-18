#include "shapes.hpp"

#include <iostream>

namespace shape
{
    double Square::LeftX() const noexcept
    {
        return centre.x - edge;
    }

    double Square::RightX() const noexcept
    {
        return centre.x + edge;
    }

    double Square::UpperY() const noexcept
    {
        return centre.y - edge;
    }

    double Square::LowerY() const noexcept
    {
        return centre.y + edge;
    }

    void Square::SetSides() noexcept
    {
        double lx = LeftX();
        double rx = RightX();
        double uy = UpperY();
        double ly = LowerY();

        sides[0].SetVectors(lx, uy, edge * 2, 0); // top side
        sides[1].SetVectors(rx, uy, 0, edge * 2); // right side
        sides[2].SetVectors(lx, uy, 0, edge * 2); // left side
        sides[3].SetVectors(lx, ly, edge * 2, 0); // bottom side
    }

    std::vector<Vector> Square::GetSides() const noexcept
    {
        return sides;
    }

    void Square::SetAnglesSides() noexcept
    {
        sides[0].setAngle();
        sides[1].setAngle();
        sides[2].setAngle();
        sides[3].setAngle();
    }

    Square::Square(Point centreCoords, double side, double alpha) noexcept :
        Shape(centreCoords, side, side, alpha),
        edge((side < 0) ? 0 : side / 2),
        sides(4, Vector())
    {
        SetDirection(alpha);
        SetSides();
        SetAnglesSides();
    }

    Square::Square(double centreX, double centreY, double side, double alpha) noexcept :
        Square(Point(centreX, centreY), side, alpha) { }

    const int Square::sideAmount() const noexcept
    {
        return 4;
    }

    Square static InitFromStdin() noexcept
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

        Square sqr = Square(Point(x, y), side, angle);
        return sqr;
    }

    void Square::Move() noexcept
    {
        centre.x += direction.first;
        centre.y -= direction.second;
        SetSides();
    }

    void Square::Draw() const noexcept
    {
        double lx = LeftX();
        double rx = RightX();
        double uy = UpperY();
        double ly = LowerY();
        al_draw_filled_rectangle(lx, uy, rx, ly, color);
    }
}
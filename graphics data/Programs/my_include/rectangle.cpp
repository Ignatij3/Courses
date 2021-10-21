#include "shapes.hpp"

#include <iostream>

namespace shape
{
    double Rectangle::LeftX() const noexcept
    {
        return centre.x - width;
    }

    double Rectangle::RightX() const noexcept
    {
        return centre.x + width;
    }

    double Rectangle::UpperY() const noexcept
    {
        return centre.y - height;
    }

    double Rectangle::LowerY() const noexcept
    {
        return centre.y + height;
    }

    void Rectangle::SetSides() noexcept
    {
        double lx = LeftX();
        double rx = RightX();
        double uy = UpperY();
        double ly = LowerY();

        sides[0].SetVectors(lx, uy, width * 2, 0);  // top side
        sides[1].SetVectors(rx, uy, 0, height * 2); // right side
        sides[2].SetVectors(lx, uy, 0, height * 2); // left side
        sides[3].SetVectors(lx, ly, width * 2, 0);  // bottom side
    }

    void Rectangle::SetSidesSetAngle() noexcept
    {
        SetSides();
        sides[0].setAngle();
        sides[1].setAngle();
        sides[2].setAngle();
        sides[3].setAngle();
    }

    Rectangle::Rectangle(Movement move, Point centreCoords, double width, double height, double alpha) noexcept :
        Shape(move, centreCoords, width, height, alpha),
        width((width < 0) ? 0 : width / 2),
        height((height < 0) ? 0 : height / 2),
        sides(4, Vector())
    {
        printf("centre coords: (%f, %f)\n", centre.x, centre.y);
        printf("width, height: %f, %f\n", width, height);

        SetDirection(alpha);
        SetSidesSetAngle();
    }

    Rectangle::Rectangle(Movement move, double centreX, double centreY, double width, double height, double alpha) noexcept :
        Rectangle(move, Point(centreX, centreY), width, height, alpha) { }

    Rectangle::Rectangle() noexcept { }

    Rectangle static InitFromStdin() noexcept
    {
        int move;
        double x, y;
        double width, height;
        double angle;

        printf("Is rectangle animate? (1/0): ");
        scanf("%d\n", move);

        printf("Enter x-coordinate of square's centre: ");
        scanf("%f\n", x);

        printf("Enter y-coordinate of square's centre: ");
        scanf("%f\n", y);

        printf("Enter rectangle's width: ");
        scanf("%f\n", width);

        printf("Enter rectangle's height: ");
        scanf("%f\n", height);

        printf("Enter square's movement angle: ");
        scanf("%f\n", angle);

        Rectangle sqr = Rectangle((move) ? shape::Movement::DYNAMIC : shape::Movement::STATIC,
            Point(x, y), width, height, angle);
        return sqr;
    }

    std::vector<Vector> Rectangle::GetSides() const noexcept
    {
        return sides;
    }

    const int Rectangle::sideAmount() const noexcept
    {
        return 4;
    }

    void Rectangle::Move() noexcept
    {
        if (dynamic)
        {
            centre.x += direction.first;
            centre.y -= direction.second;
            SetSides();
        }
    }

    void Rectangle::Draw() const noexcept
    {
        double lx = LeftX();
        double rx = RightX();
        double uy = UpperY();
        double ly = LowerY();
        al_draw_filled_rectangle(lx, uy, rx, ly, color);
    }
}
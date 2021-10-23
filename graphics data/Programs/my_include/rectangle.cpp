#include "shapes.hpp"

#include <iostream>

namespace shape
{
    long double Rectangle::LeftX() const noexcept
    {
        return centre.x - width;
    }

    long double Rectangle::RightX() const noexcept
    {
        return centre.x + width;
    }

    long double Rectangle::UpperY() const noexcept
    {
        return centre.y - height;
    }

    long double Rectangle::LowerY() const noexcept
    {
        return centre.y + height;
    }

    void Rectangle::SetSides() noexcept
    {
        long double lx = LeftX();
        long double rx = RightX();
        long double uy = UpperY();
        long double ly = LowerY();

        sides[0]->SetVectors(lx, uy, width * 2, 0);  // top side
        sides[1]->SetVectors(rx, uy, 0, height * 2); // right side
        sides[2]->SetVectors(lx, uy, 0, height * 2); // left side
        sides[3]->SetVectors(lx, ly, width * 2, 0);  // bottom side
    }

    void Rectangle::SetSidesSetAngle() noexcept
    {
        SetSides();
        sides[0]->setAngle();
        sides[1]->setAngle();
        sides[2]->setAngle();
        sides[3]->setAngle();
    }

    Rectangle::Rectangle(Movement move, Point centreCoords, double width, double height, double alpha) :
        Shape(move, centreCoords, width, height, alpha),
        width((width < 0) ? 0 : width / 2),
        height((height < 0) ? 0 : height / 2)
    {
        for (int i = 0; i < SIDE_AMOUNT; i++)
            sides.insert(sides.end(), new Vector());

        SetDirection();
        SetSidesSetAngle();
    }

    Rectangle::Rectangle(Movement move, long double centreX, long double centreY, double width, double height, double alpha) :
        Rectangle(move, Point(centreX, centreY), width, height, alpha) { }

    Rectangle::Rectangle() noexcept { }

    Rectangle::~Rectangle() //needs fix
    {
        // delete sides[0];
        // delete sides[1];
        // delete sides[2];
        // delete sides[3];
    }

    Rectangle static InitFromStdin() noexcept
    {
        int move;
        long double x, y;
        double width, height;
        double angle;

        printf("Is rectangle animate? (1/0): ");
        scanf("%d\n", move);

        printf("Enter x-coordinate of square's centre: ");
        scanf("%Lf\n", x);

        printf("Enter y-coordinate of square's centre: ");
        scanf("%Lf\n", y);

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

    std::vector<Vector*> Rectangle::GetSides() const noexcept
    {
        return sides;
    }

    const int Rectangle::sideAmount() const noexcept
    {
        return SIDE_AMOUNT;
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
        long double lx = LeftX();
        long double rx = RightX();
        long double uy = UpperY();
        long double ly = LowerY();
        al_draw_filled_rectangle(lx, uy, rx, ly, color);
    }

    Rectangle& Rectangle::operator=(const Rectangle& rhs) noexcept
    {
        angle   = rhs.angle;
        centre  = rhs.centre;
        color   = rhs.color;
        dynamic = rhs.dynamic;
        height  = rhs.height;
        width   = rhs.width;
        sides   = rhs.sides;

        return *this;
    }
}
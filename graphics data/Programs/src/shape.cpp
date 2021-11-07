#include "screen.hpp"
#include "shapes.hpp"

#include <cmath>
#include <iostream>

// move in other header
bool AlmostEqual(double a, double b, double epsilon)
{
    return std::abs(b - a) <= epsilon;
}

namespace shape
{
    Shape::Shape(Movement move, Point centreCoords, double width, double height, double alpha) noexcept :
        centre(centreCoords),
        angle((alpha < 0) ? 0
                          : ((alpha > 360) ? 360 : alpha)),
        sides(1, nullptr)
    {
        double half_width  = width / 2;
        double half_height = height / 2;

        centre.x = (centre.x < half_width) ? half_width
                                           : ((centre.x > (screen::Window::window_width - half_width)) ? (screen::Window::window_width - half_width) : centre.x);
        centre.y = (centre.y < half_height) ? half_height
                                            : ((centre.y > (screen::Window::window_height - half_height)) ? (screen::Window::window_height - half_height) : centre.y);

        centre = screen::ConvertToNormalCoords(centre);
        MovementToggle(move);
    }

    Shape::Shape() noexcept { }

    void Shape::Reflect(double otherVectorAngle) noexcept
    {
        if (dynamic)
        {
            angle = -360 + (2 * otherVectorAngle) - angle;
            angle = (angle < 0) ? 360 + angle : angle;
            SetDirection();
        }
    }

    // setDirection sets shape's direction in degrees, where 0 points right and goes anticlockwise
    // if alpha is smaller than 0 or greater than 360, direction is set to 0
    void Shape::SetDirection() noexcept
    {
        direction.first  = cos(angle * PI / 180); // converting degrees to radians
        direction.second = sin(angle * PI / 180); // converting degrees to radians
    }

    std::pair<bool, std::pair<const Vector*, const Vector*>> Shape::CollideWith(Shape* other) noexcept
    {
        std::vector<Vector> normals       = GetNormals();
        std::vector<Vector> other_normals = other->GetNormals();
    }

    void Shape::MovementToggle() noexcept
    {
        dynamic = dynamic ? false : true;
    }

    void Shape::MovementToggle(Movement move) noexcept
    {
        dynamic = move;
    }
}
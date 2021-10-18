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
    Shape::Shape(Point centreCoords, double width, double height, double alpha) noexcept :
        centre(centreCoords),
        angle((alpha < 0) ? 0
                          : ((alpha > 360) ? 360 : alpha)),
        sides(1, Vector())
    {
        // if window initialised
        centre.x = (centre.x < width) ? width
                                      : ((centre.x > (screen::Window::window_width - width)) ? (screen::Window::window_width - width) : centre.x);
        centre.y = (centre.y < height) ? height
                                       : ((centre.y > (screen::Window::window_height - height)) ? (screen::Window::window_height - height) : centre.y);

        centre = screen::ConvertToNormalCoords(centre);
    }

    void Shape::Reflect(double otherVectorAngle) noexcept
    {
        double newAngle = -360 + (2 * otherVectorAngle) - angle;
        newAngle        = (newAngle < 0) ? 360 + newAngle : newAngle;
        SetDirection(newAngle);
    }

    // setDirection sets shape's direction in degrees, where 0 points right and goes anticlockwise
    // if alpha is smaller than 0 or greater than 360, direction is set to 0
    void Shape::SetDirection(double alpha) noexcept
    {
        angle = (alpha < 0 || alpha > 360) ? 0 : alpha;

        direction.first  = cos(angle * PI / 180); // converting degrees to radians
        direction.second = sin(angle * PI / 180); // converting degrees to radians
    }

    std::pair<bool, std::pair<const Vector, const Vector>> Shape::CollideWith(const Shape* other) const noexcept
    {
        std::vector<Vector> sides       = GetSides();
        std::vector<Vector> other_sides = other->GetSides();

        for (int i = 0; i < sideAmount(); ++i) //проверять с какой стороной реально столкнулись (проверять разницу между координатами x и y
                                               //если 2 стороны на одинаковом расстоянии, значит взять перпендикуляр биссектрисы угла)
        {
            for (int j = 0; j < other->sideAmount(); ++j)
            {
                if (sides[i].Cross(other_sides[j]))
                    return std::make_pair(true, FindSidesToReflect(sides, other_sides, i, j));
            }
        }

        return std::make_pair(false, std::make_pair(Vector(), Vector()));
    }

    std::pair<bool, const Vector> Shape::LiesOnLine(const std::vector<Vector>& sides, const Point& angle) const noexcept
    {

        std::size_t sideAmt = sides.size();
        for (int i = 0; i < sideAmt; ++i) // find vector which is touched by angle
        {
            Vector otherVec = sides[i];
            Vector temp(otherVec.a, angle);

            if (sides[i].a == angle)                   // angle touches angle of line
                return std::make_pair(true, otherVec); // vector, perpendicular to incoming angle

            if (AlmostEqual(otherVec.Slope(), temp.Slope(), 1) && temp.LiesBetween(otherVec)) // lies on the same line as the vector does and is between endpoints
                return std::make_pair(true, otherVec);
        }

        return std::make_pair(false, Vector());
    }

    std::pair<const Vector, const Vector> Shape::FindSidesToReflect(std::vector<Vector>& shapeSides, std::vector<Vector>& otherShapeSides, int sideIndex, int otherSideIndex) const noexcept
    {
        std::pair<Vector, Vector> resulting_vectors; // first vector - side of first shape, second - of second shape

        if (std::pair<bool, const Vector> res = LiesOnLine(otherShapeSides, shapeSides[sideIndex].a); res.first) // check whether and which side does first angle of first collided vector touch
            resulting_vectors.second = res.second;
        else if (std::pair<bool, const Vector> res = LiesOnLine(otherShapeSides, shapeSides[sideIndex].b); res.first) // second angle of first vector
            resulting_vectors.second = res.second;
        else // side touched purely angle
            resulting_vectors.second = shapeSides[sideIndex];

        if (std::pair<bool, const Vector> res = LiesOnLine(shapeSides, otherShapeSides[otherSideIndex].a); res.first) // first angle of second vector
            resulting_vectors.first = res.second;
        else if (std::pair<bool, const Vector> res = LiesOnLine(shapeSides, otherShapeSides[otherSideIndex].b); res.first) // second angle of second vector
            resulting_vectors.first = res.second;
        else
            resulting_vectors.first = otherShapeSides[otherSideIndex];

        return resulting_vectors;
    }

    std::pair<bool, const Vector> Shape::CollideWith(const Vector* other_vector) const noexcept
    {

        std::vector<Vector> derived_sides = GetSides();

        for (int i = 0; i < sideAmount(); ++i)
        {
            if (derived_sides[i].Cross(*other_vector))
                return std::make_pair(true, *other_vector);
        }

        return std::make_pair(false, Vector());
    }
}
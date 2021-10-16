#include "shapes.hpp"

#include <cmath>
#include <iostream>

namespace shape
{
    Shape::Shape(Point<double> centreCoords, double alpha) :
        centre(centreCoords),
        angle((alpha < 0) ? 0
                          : ((alpha > 360) ? 360 : alpha)),
        sides(1, Vector()) { }

    void Shape::Reflect(double otherVectorAngle)
    {
        double newAngle = -360 + (2 * otherVectorAngle) - angle;
        newAngle        = (newAngle < 0) ? 360 + newAngle : newAngle;
        SetDirection(newAngle);
    }

    // setDirection sets shape's direction in degrees, where 0 points right and goes anticlockwise
    // if alpha is smaller than 0 or greater than 360, direction is set to 0
    void Shape::SetDirection(double alpha)
    {
        angle = (alpha < 0 || alpha > 360) ? 0 : alpha;
        printf("new angle: %f\n", angle);

        direction.first  = cos(angle * PI / 180); // converting degrees to radians
        direction.second = sin(angle * PI / 180); // converting degrees to radians
    }

    std::pair<bool, std::pair<const Vector, const Vector>> Shape::getVectorIfCollide(const Shape* other) const
    {
        std::vector<Vector> sides       = GetSides();
        std::vector<Vector> other_sides = other->GetSides();

        for (int i = 0; i < sideAmount(); ++i) //проверять с какой стороной реально столкнулись (проверять разницу между координатами x и y
                                               //если 2 стороны на одинаковом расстоянии, значит взять перпендикуляр биссектрисы угла)
        {
            for (int j = 0; j < other->sideAmount(); ++j)
            {
                if (sides[i].Cross(other_sides[j]))
                    return std::make_pair(true, std::make_pair(sides[i], other_sides[j])); // returns horisontal line
            }
        }

        return std::make_pair(false, std::make_pair(sides[0], other_sides[0]));
    }

    std::pair<bool, const Vector> Shape::getVectorIfCollide(const Vector* other_vector) const
    {

        std::vector<Vector> derived_sides = GetSides();

        for (int i = 0; i < sideAmount(); ++i)
        {
            if (derived_sides[i].Cross(*other_vector))
                return std::make_pair(true, *other_vector);
        }

        return std::make_pair(false, derived_sides[0]);
    }
}
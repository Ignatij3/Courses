#include "shapes.hpp"

#include <cmath>

namespace shape
{
    Shape::Shape(Point<double> centreCoords, double alpha) :
        centre(centreCoords),
        angle((alpha < 0) ? 0
                          : ((alpha > 360) ? 360 : alpha)) { }

    void Shape::Reflect(double otherVectorAngle)
    {
        double newAngle = -360 + (2 * otherVectorAngle) - angle;
        newAngle        = (newAngle < 0) ? 360 + newAngle : newAngle;
        SetDirection(newAngle);
    }

    //setDirection sets shape's direction in degrees, where 0 points right and goes anticlockwise
    //if alpha is smaller than 0 or greater than 360, direction is set to 0
    void Shape::SetDirection(double alpha)
    {
        angle = (alpha < 0 || alpha > 360) ? 0 : alpha;

        direction.first  = cos(angle * PI / 180); //converting degrees to radians
        direction.second = sin(angle * PI / 180); //converting degrees to radians
    }

    std::pair<bool, const Vector&> Shape::getVectorIfCollide(const Shape* other) const
    {
        for (int i = 0; i < sideAmount(); ++i)
        {
            for (int j = 0; j < other->sideAmount(); ++j)
            {
                if (sides[i].Cross(other->sides[j]))
                    return std::make_pair(true, other->sides[j]);
            }
        }

        return std::make_pair(false, sides[0]);
    }
}
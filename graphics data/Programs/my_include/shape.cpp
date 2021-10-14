#include "shapes.hpp"

#include <cmath>
#include <iostream>

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

    std::pair<bool, const Vector> Shape::getVectorIfCollide(const Shape* other) const
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

    std::pair<bool, const Vector*> Shape::getVectorIfCollide(const Vector* other_vector) const
    {

        // printf("COLLISION side[0]: %.1f %.1f %.1f %.1f\n", sides[0].a.x, sides[0].a.y, sides[0].a.x + sides[0].b.x, sides[0].a.y + sides[0].b.y);
        // printf("COLLISION side[1]: %.1f %.1f %.1f %.1f\n", sides[1].a.x, sides[1].a.y, sides[1].a.x + sides[1].b.x, sides[1].a.y + sides[1].b.y);
        // printf("COLLISION side[2]: %.1f %.1f %.1f %.1f\n", sides[2].a.x, sides[2].a.y, sides[2].a.x + sides[2].b.x, sides[2].a.y + sides[2].b.y);
        // printf("COLLISION side[3]: %.1f %.1f %.1f %.1f\n\n", sides[3].a.x, sides[3].a.y, sides[3].a.x + sides[3].b.x, sides[3].a.y + sides[3].b.y);

        for (int i = 0; i < sideAmount(); ++i)
        {
            if (sides[i].Cross(*other_vector))
                return std::make_pair(true, other_vector);
        }

        return std::make_pair(false, &sides[0]);
    }
}
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

        direction.first  = cos(angle * PI / 180); // converting degrees to radians
        direction.second = sin(angle * PI / 180); // converting degrees to radians
    }

    std::pair<bool, std::pair<const Vector, const Vector>> Shape::CollideWith(const Shape* other) const
    {
        std::vector<Vector> sides       = GetSides();
        std::vector<Vector> other_sides = other->GetSides();

        for (int i = 0; i < sideAmount(); ++i) //проверять с какой стороной реально столкнулись (проверять разницу между координатами x и y
                                               //если 2 стороны на одинаковом расстоянии, значит взять перпендикуляр биссектрисы угла)
        {
            for (int j = 0; j < other->sideAmount(); ++j)
            {
                if (sides[i].Cross(other_sides[j]))
                    // return std::make_pair(true, std::make_pair(sides[i], other_sides[j])); // returns horisontal line
                    return std::make_pair(true, FindSidesToReflect(sides, other_sides, i, j));
            }
        }

        return std::make_pair(false, std::make_pair(sides[0], other_sides[0]));
    }

    std::pair<bool, const Vector> Shape::LiesOnLine(const std::vector<Vector>& sides, const Point<double>& angle) const
    {
        std::size_t sideAmt = sides.size();
        for (int i = 0; i < sideAmt; ++i) // find vector which is touched by angle
        {
            Vector otherVec = sides[i];
            Vector temp(otherVec.a, angle);

            if (otherVec.Slope() == temp.Slope()) // lies on the same line as the vector does
            {
                double lowY  = otherVec.LowestY();
                double highY = otherVec.HighestY();
                double lowX  = otherVec.LowestX();
                double highX = otherVec.HighestX();

                if ((highX >= temp.b.x && temp.b.x >= lowX) && (highY >= temp.b.y && temp.b.y >= lowY)) // point is in between vector's endpoints, aka lies on it
                    return std::make_pair(true, otherVec);
            }
        }

        return std::make_pair(false, Vector());
    }

    std::pair<const Vector, const Vector> Shape::FindSidesToReflect(std::vector<Vector>& shapeSides, std::vector<Vector>& otherShapeSides, int sideIndex, int otherSideIndex) const
    {
        std::pair<Vector, Vector> resulting_vectors; // first vector - side of first shape, second - of second shape

        // if both
        if (std::pair<bool, const Vector> res = LiesOnLine(otherShapeSides, shapeSides[sideIndex].a); res.first) // check whether and which side does first angle of first collided vector touch
        {
            resulting_vectors.second = res.second;
            printf("1.1\n");
        }
        else if (std::pair<bool, const Vector> res = LiesOnLine(otherShapeSides, shapeSides[sideIndex].b); res.first) // second angle of first vector
        {
            resulting_vectors.second = res.second;
            printf("1.2\n");
        }
        else // side touched purely angle (or mistake in code)
        {
            resulting_vectors.second = shapeSides[sideIndex];
            printf("1.0\n");
        }

        // if both
        if (std::pair<bool, const Vector> res = LiesOnLine(shapeSides, otherShapeSides[otherSideIndex].a); res.first) // first angle of second vector
        {
            resulting_vectors.first = res.second;
            printf("2.1\n");
        }
        else if (std::pair<bool, const Vector> res = LiesOnLine(shapeSides, otherShapeSides[otherSideIndex].b); res.first) // second angle of second vector
        {
            resulting_vectors.first = res.second;
            printf("2.2\n");
        }
        else
        {
            resulting_vectors.first = otherShapeSides[otherSideIndex];
            printf("2.0\n");
        }

        printf("vector angles: %.2f, %.2f\n", resulting_vectors.first.getAngle(), resulting_vectors.second.getAngle());
        return resulting_vectors;
    }

    std::pair<bool, const Vector> Shape::CollideWith(const Vector* other_vector) const
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
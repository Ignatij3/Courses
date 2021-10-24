#include "shapes.hpp"

#include <iostream>

namespace shape
{

    Square::Square(Movement move, Point centreCoords, double side, double alpha) noexcept :
        Rectangle(move, centreCoords, side, side, alpha) { }

    Square::Square(Movement move, long double centreX, long double centreY, double side, double alpha) noexcept :
        Rectangle(move, Point(centreX, centreY), side, side, alpha) { }

    Square static InitFromStdin() noexcept
    {
        int move;
        long double x, y;
        double side;
        double angle;

        printf("Is square animate? (1/0): ");
        scanf("%d\n", move);

        printf("Enter x-coordinate of square's centre: ");
        scanf("L%f\n", x);

        printf("Enter y-coordinate of square's centre: ");
        scanf("%Lf\n", y);

        printf("Enter square's side: ");
        scanf("%f\n", side);

        printf("Enter square's movement angle: ");
        scanf("%f\n", angle);

        Square sqr = Square((move) ? shape::Movement::DYNAMIC : shape::Movement::STATIC,
            Point(x, y), side, angle);
        return sqr;
    }
}
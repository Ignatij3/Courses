#include "figure.hpp"

#include <array>
#include <iostream>
#include <numeric>

int main()
{
    Square sqr(15);
    Rectangle rect(7, 3);
    Circle circ(1);

    std::array<Figure*, 3> arr;
    arr[0] = &sqr;
    arr[1] = &rect;
    arr[2] = &circ;

    std::cout << std::accumulate(arr.begin(), arr.end(), 0.0, Figure::totalArea);
}
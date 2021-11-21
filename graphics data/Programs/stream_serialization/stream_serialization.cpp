#include "factory.hpp"
#include "figures.hpp"

#include <fstream>
#include <iostream>
#include <sstream>

std::fstream& operator<<(std::fstream& os, const Figure* fig)
{
    os << fig->toFileString();
    return os;
}

void writeToFile(const Figure* fig)
{
    std::fstream fout("figures.dat", std::ios::binary | std::ios::out | std::ios::app);
    fout << fig;
    fout.close();
}

std::string getFromFile()
{
    std::fstream fin("figures.dat", std::ios::binary | std::ios::in);
    std::stringstream buffer;
    buffer << fin.rdbuf();
    fin.close();

    return buffer.str();
}

int main()
{
    FigureFactory fig("Square:side=200,x=-13.333333,y=0;Circle:x=15,y=15,radius=15;");
    fig.addFigures(getFromFile());

    Figure* ptr = static_cast<Figure*>(fig.getNext());
    do
    {
        printf("%s\n", ptr->toString().c_str());
        ptr = static_cast<Figure*>(fig.getNext());
        // writeToFile(ptr);
    } while (ptr);

    return 0;
}
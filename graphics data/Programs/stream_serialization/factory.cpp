#include "factory.hpp"

#include "figures.hpp"

#include <iostream>

FigureFactory::FigureFactory() { }

FigureFactory::FigureFactory(const std::string& str) :
    figureData(str, std::ios::app | std::ios::in | std::ios::out) { }

void FigureFactory::addFigures(const std::string& str)
{
    figureData << str;
}

void* FigureFactory::getNext()
{
    std::string curr_str;
    std::string class_name;

    std::getline(figureData, curr_str, ';');
    if (curr_str == "")
        return nullptr;

    std::getline(std::stringstream(curr_str), class_name, ':');
    if (class_name == curr_str) //taken from file
        return getNextStream(std::stringstream(curr_str));

    Figure* figure_ptr = nullptr;

    if (class_name == "Square")
    {
        figure_ptr = new Square;
        figure_ptr->fromString(curr_str.substr(7));
    }
    else if (class_name == "Circle")
    {
        figure_ptr = new Circle;
        figure_ptr->fromString(curr_str.substr(7));
    }
    else
        std::printf("Object cannot be parsed\n");

    return figure_ptr;
}

void* FigureFactory::getNextStream(std::stringstream str_stream)
{
    std::string class_name;
    Figure* figure_ptr = nullptr;

    std::getline(str_stream, class_name, ' ');
    if (class_name == "Square")
    {
        figure_ptr = new Square;
        figure_ptr->fromString(str_stream);
    }
    else if (class_name == "Circle")
    {
        figure_ptr = new Circle;
        figure_ptr->fromString(str_stream);
    }
    else
        std::printf("Object cannot be parsed\n");

    return figure_ptr;
}
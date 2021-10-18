#ifndef GRAPHICS_WINDOW_H_
#define GRAPHICS_WINDOW_H_

#include "shapes.hpp"

#include <allegro5/allegro.h>
#include <allegro5/allegro_primitives.h>
#include <exception>
#include <string>

namespace screen
{

    inline shape::Vector walls[4]; // make stationary rectangle

    std::pair<double, double> ConvertToNormalCoords(double x, double y);
    shape::Point ConvertToNormalCoords(shape::Point a);

    class Window {
      protected:
        struct InitFailure : public std::exception {
            InitFailure(const std::string& msg);

            const char* what() const throw();

          protected:
            std::string message;
        };

        ALLEGRO_DISPLAY* alDisplay;
        ALLEGRO_EVENT_QUEUE* alEventQueue;
        ALLEGRO_TIMER* alTimer;
        std::pair<std::vector<shape::Vector>, std::vector<shape::Shape*>> objects;
        bool exit = false;

        void InitAllegro();
        void DestroyAllegro();
        void RunAllegro(void (*fpsCallback)(), void (*drawCallback)());

      public:
        static int fps;
        static int window_width; // make read only
        static int window_height;

        Window(int fps_, int win_width, int win_height);

        void Run(void (*fpsCallback)(), void (*drawCallback)());
        void AddObject(shape::Vector) noexcept;
        void AddObject(shape::Shape& figure) noexcept;
        void SetColor(shape::Shape& figure, ALLEGRO_COLOR col) noexcept;
        void AddObject(shape::Shape& figure, ALLEGRO_COLOR col) noexcept;
        void ManageCollisions() const noexcept;
        void MoveAll() noexcept;
        void DrawAll() const noexcept;
    };
}

#endif // !GRAPHICS_WINDOW_H_
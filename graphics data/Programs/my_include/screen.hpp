#ifndef GRAPHICS_WINDOW_H_
#define GRAPHICS_WINDOW_H_

#include "shapes.hpp"

#include <exception>
#include <string>

#define BLACK     al_map_rgb(0, 0, 0)
#define RED       al_map_rgb(255, 0, 0)
#define GREEN     al_map_rgb(0, 255, 0)
#define BLUE      al_map_rgb(0, 0, 255)
#define DARKRED   al_map_rgb(128, 0, 0)
#define DARKGREEN al_map_rgb(0, 128, 0)
#define DARKBLUE  al_map_rgb(0, 0, 128)
#define PINK      al_map_rgb(255, 20, 147)
#define ORANGE    al_map_rgb(255, 140, 0)
#define YELLOW    al_map_rgb(255, 255, 0)
#define PURPLE    al_map_rgb(147, 112, 219)
#define BROWN     al_map_rgb(160, 82, 45)
#define BEIGE     al_map_rgb(210, 180, 140)
#define LIGHTGRAY al_map_rgb(211, 211, 211)
#define DARKGRAY  al_map_rgb(105, 105, 105)
#define WHITE     al_map_rgb(250, 250, 250)

typedef void (*FpsCallback)();
typedef void (*DrawCallback)();

namespace screen
{
    inline shape::Vector walls[4]; // make stationary rectangle

    std::pair<double, double> ConvertToNormalCoords(double x, double y);
    shape::Point ConvertToNormalCoords(shape::Point a);

    class Window {
      private:
        struct InitFailure : public std::exception {
            InitFailure(const std::string& msg);

            const char* what() const throw();

          protected:
            std::string message;
        };

        std::vector<shape::Shape*> objects;
        shape::Rectangle window_frame;

        bool initialized = false;
        bool exit        = false;

        ALLEGRO_DISPLAY* alDisplay        = nullptr;
        ALLEGRO_EVENT_QUEUE* alEventQueue = nullptr;
        ALLEGRO_TIMER* alTimer            = nullptr;

        void InitAllegro();
        void DestroyAllegro();
        void RunAllegro(FpsCallback fpsCallback, DrawCallback drawCallback);

      public:
        static int fps;
        static int window_width; // make read only
        static int window_height;

        Window(int fps_, int win_width, int win_height);

        void Run(FpsCallback fpsCallback, DrawCallback drawCallback);
        void AddObject(shape::Shape& figure);
        void AddObject(shape::Shape& figure, ALLEGRO_COLOR col);
        void SetColor(shape::Shape& figure, ALLEGRO_COLOR col) noexcept;
        void ManageCollisions() const noexcept;
        void MoveAll() noexcept;
        void DrawAll() const noexcept;

        explicit operator bool() const;
    };
}

#endif // !GRAPHICS_WINDOW_H_
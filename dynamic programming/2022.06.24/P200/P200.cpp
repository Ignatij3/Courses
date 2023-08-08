#include <iostream>
#include <stack>
#include <vector>

struct node;
std::vector<std::vector<std::string>> get_data();
std::string calculate(std::vector<std::vector<std::string>>);
void trim_strings(std::vector<std::vector<std::string>>& index);
node* find_node(std::vector<node*>, char);
bool is_connection(node*, node*);
std::string topological_sort(std::vector<node*>& graph);
void dfs(std::vector<node*>& graph, node* elem, std::stack<node*>& stck);

struct node {
    std::vector<node*> next_letters;
    char letter;
    bool visited = false;
    node(char letter);
};

node::node(char letter)
{
    this->letter = letter;
}

std::vector<std::vector<std::string>> get_data()
{
    std::string word;
    std::vector<std::string> index;

    while (word != "#") {
        getline(std::cin, word);
        index.push_back(word);
    }

    index.pop_back();

    std::vector<std::vector<std::string>> res;
    res.push_back(index);
    return res;
}

int main()
{
    std::vector<std::vector<std::string>> index = get_data();

    std::string res = calculate(index);
    std::cout << res << std::endl;
}

std::string calculate(std::vector<std::vector<std::string>> index)
{
    std::vector<node*> letters;
    char old_letter = '\0';
    char letter;

    while (index.size() > 0) {
        for (int i = 0; i < index.size(); i++) {
            for (int j = 0; j < index[i].size(); j++) {
                letter = index[i][j][0];

                if (letter != old_letter && find_node(letters, letter) == nullptr) {
                    node* new_letter_node = new node(letter);
                    letters.push_back(new_letter_node);

                    if (old_letter != '\0') {
                        node* old_letter_node = find_node(letters, old_letter);
                        old_letter_node->next_letters.push_back(new_letter_node);
                    }
                }

                old_letter = letter;
            }

            trim_strings(index);
        }
    }

    std::string res = topological_sort(letters);

    for (int i = 0; i < letters.size(); i++)
        delete letters[i];

    return res;
}

void trim_strings(std::vector<std::vector<std::string>>& index)
{
    int counter;
    char old_letter;
    char curr_letter;
    std::vector<std::vector<std::string>> new_index;

    for (int i = 0; i < index.size(); i++) {
        std::vector<std::string> str_block;
        counter    = -1;
        old_letter = index[i][0][0];

        for (int j = 0; j < index[i].size(); j++) {
            curr_letter = index[i][j][0];
            index[i][j].erase(0, 1);

            if (curr_letter != old_letter) {
                if (counter != 0 && str_block.size() > 0)
                    new_index.push_back(str_block);

                str_block.clear();
                counter = -1;
            }

            if (index[i][j].size() > 0)
                str_block.push_back(index[i][j]);
            old_letter = curr_letter;
            counter++;
        }

        if (counter != 0 && str_block.size() > 0)
            new_index.push_back(str_block);
    }

    index = new_index;
}

node* find_node(std::vector<node*> letters, char letter)
{
    for (int i = 0; i < letters.size(); i++) {
        if (letters[i]->letter == letter) {
            return letters[i];
        }
    }

    return nullptr;
}

bool is_connection(node* char1, node* char2)
{
    for (int i = 0; i < char1->next_letters.size(); i++) {
        if (char1->next_letters[i] == char2) {
            return true;
        }
    }

    return false;
}

std::string topological_sort(std::vector<node*>& graph)
{
    std::string res;
    std::stack<node*> stck;

    for (int i = 0; i < graph.size(); i++)
        if (!graph[i]->visited)
            dfs(graph, graph[i], stck);

    while (!stck.empty()) {
        res += stck.top()->letter;
        stck.pop();
    }

    return res;
}

void dfs(std::vector<node*>& graph, node* elem, std::stack<node*>& stck)
{
    elem->visited = true;
    std::vector<node*>::iterator i;
    for (int i = 0; i < elem->next_letters.size(); i++)
        if (!elem->next_letters[i]->visited)
            dfs(graph, elem->next_letters[i], stck);

    stck.push(elem);
}

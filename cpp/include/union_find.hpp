#ifndef __PXU_UNION_FIND_HPP__
#define __PXU_UNION_FIND_HPP__

#include <memory>
#include <unordered_map>

class UnionFind {
public:
    UnionFind() =default;
    ~UnionFind() =default;

    UnionFind(const UnionFind&) =delete;
    UnionFind& operator=(const UnionFind&) =delete;
    UnionFind(UnionFind&&) =delete;
    UnionFind& operator=(UnionFind&&) =delete;

    void add(int v);
    int find(int v);
    void group(int a, int b);

protected:
    class Node {
    public:
        explicit Node(int v) { value = v; }
        Node() =default;
        ~Node() =default;
        
        Node(const Node&) =delete;
        Node& operator=(const Node&) =delete;
        Node(Node&&) =delete;
        Node& operator=(Node&&) =delete;

        int value;
        int rank = 0;
        std::shared_ptr<Node> root = nullptr;
    };

    std::shared_ptr<Node> _find(std::shared_ptr<Node> node);
    void _group(std::shared_ptr<Node> a, std::shared_ptr<Node> b);

    std::unordered_map<int, std::shared_ptr<Node>> m_nodeMap;
};

#endif
#ifndef PROCESSFILE_H
#define PROCESSFILE_H

#include <vector>

std::vector<std::vector<float>> readFile(const char * filePath);
bool checkMatrixFormat(std::vector<std::vector<float>>& T, int n, int m);
#endif
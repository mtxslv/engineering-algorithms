#ifndef PROCESSFILE_H
#define PROCESSFILE_H

#include <vector>

std::vector<std::vector<float>> readFile(const char * filePath);
bool checkMatrixSquare(std::vector<std::vector<float>>& T);
#endif
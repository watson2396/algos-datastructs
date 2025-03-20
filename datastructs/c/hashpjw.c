#include "hashpjw.h"

unsigned int hashpjw(const void* key)
{

    const char* ptr;
    unsigned int val;
    unsigned int PRIME_TBLSIZ;

    // The number 3 is almost certainly not right
    // fix with more realistic number later
    PRIME_TBLSIZ = 3;

    //  Hash the key by performing a number of bit operations on it

    val = 0;
    ptr = key;

    while (*ptr != '\0')
    {

        unsigned int tmp;

        val = (val << 4) + (*ptr);

        if (tmp = (val & 0xf0000000))
        {

            val = val ^ (tmp >> 24);
            val = val ^ tmp;

        }

        ptr++;

    }

    //  In practice, replace PRIME_TBLSIZ with the actual table size

    return val % PRIME_TBLSIZ;

}
# 1.8 Input Management (IM) Interface Information

Note: This section is subject to change in support of transitioning away
from IM.

+----------------------+-------------------------+---------------------+
| **Path**             | **Value**               | **Notes**           |
+======================+=========================+=====================+
| Original Filename    | FROXK.Agency.SPR.Unique | Dataset name        |
|                      |                         | including           |
|                      | Unique=Specified by the | delimiters is up to |
|                      | agency to make the      | 44 characters long. |
|                      | dataset unique for the  | Each node can       |
|                      | day.                    | contain up to 8     |
|                      |                         | characters.         |
|                      |                         |                     |
|                      |                         | If the same dataset |
|                      |                         | name is used within |
|                      |                         | a day, the previous |
|                      |                         | dataset will be     |
|                      |                         | overwritten.        |
+----------------------+-------------------------+---------------------+
| ControlNumber        | Cnnnnnn                 | Assigned by IM      |
+----------------------+-------------------------+---------------------+
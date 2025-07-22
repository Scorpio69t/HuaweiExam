#include <conio.h>
#include <ctype.h>
#include <math.h>
#include <memory.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <windows.h>

#define MAX 999
#define MAXN 9

typedef enum bool { false, true } bool; // 定义布尔类型的枚举

// 玩家信息结点
typedef struct _player {
  int m;                // 所用分钟数
  int s;                // 所用秒数
  char name[20];        // 玩家姓名
  int level;            // 游戏难度
  struct _player *next; // 指向下一个玩家结点
} player;

void pause(const char *str, ...); // 暂停程序

void show(player *easy, player *normal, player *hard); // 显示排名情况

void order(player *head); // 按所用时间从少到多进行排序

player *get_record(int level); // 获取排名记录

void record(player info); // 记录玩家的游戏时间

bool judge(int *player_res, int *answer); // 判断是否回答正确

void ready(); // 给用户5秒的观察时间

int get_time(); // 获取时间

bool receiver(int *player_res); // 获取用户输入

void sudoku_level(int *answer, int count); // 根据难度初始化初盘

void print(int *answer); // 打印数独

void showHelp(); // 显示帮助菜单

char printMainMenu(); // 显示菜单

bool set(int x, int y, int val);

void reset(int x, int y);

void initXOrd(int *xOrd); // 0~9随机序列

bool fillFrom(int y, int val);

void initShudu();

void get_answer(int *answer);

int row_size = 593;   // 行数
int col_size = 324;   // 列数
int result[81];       // 存放结果行的栈
int index = 0;        // 栈指针
int sudoku[81] = {0}; // 存放数独
int time_start = 0;   // 开始时间
int time_end = 0;     // 结束时间

int sudo[MAXN][MAXN]; // sudo最终盘

void main() {
  int player_res[81] = {0};
  int choice;
  int **matrix;         // 存放数独的01矩阵
  int answer[81] = {0}; // 存放答案
  int option;           // 难度选项
  char menuID;          // 菜单id
  player *easy;         // 容易难度排行
  player *normal;       // 简单难度排行
  player *hard;         // 困难难度排行
  player info;          // 玩家信息

  srand(time(NULL));
  while (true) {
    initShudu();
    get_answer(answer);
    menuID = printMainMenu(); // 显示菜单
    switch (menuID) {
    case '1':
      printf("玩家名:");
      scanf("%s", &info.name);
      if (strlen(info.name) > 20) {
        printf("名字太长！\n");
        break;
      }
      printf("请选择游戏难度:  1.简单\t2.一般\t3.困难\n");
      scanf("%d", &option);
      printf("\n");
      switch (option) {
      case 1:
        sudoku_level(answer, 75); // 挖空答案，生成初盘
        ready();
        time_start = get_time();
        if (!receiver(player_res)) {
          printf("\n您已放弃作答!\t正确答案为:\n\n");
          print(answer);
          break;
        }
        time_end = get_time();
        info.m = (time_end - time_start) / 60;
        info.s = (time_end - time_start) % 60;
        info.level = 1;

        if (judge(player_res, answer)) {
          pause("恭喜你成功了!\t用时:  %d:%d\n", info.m, info.s);
          record(info);
        } else {
          printf("\n回答错误!\t正确答案为:\n\n");
          fflush(stdin);
          print(answer);
          pause("按任意键返回...");
        }
        break;
      case 2:
        sudoku_level(answer, 35);
        time_start = get_time();
        ready();
        if (!receiver(player_res)) {
          printf("\n您已放弃作答!\t正确答案为:\n\n");
          print(answer);
          break;
        }
        time_end = get_time();
        info.m = (time_end - time_start) / 60;
        info.s = (time_end - time_start) % 60;
        info.level = 2;
        if (judge(player_res, answer)) {
          printf("恭喜你成功了!\t用时:  %d:%d\n", info.m, info.s);
          record(info);
        } else {
          printf("回答错误!\t正确答案为:\n\n");
          fflush(stdin);
          print(answer);
        }
        break;
      case 3:
        sudoku_level(answer, 30);
        time_start = get_time();
        ready();
        if (!receiver(player_res)) {
          printf("\n您已放弃作答!\t正确答案为:\n\n");
          print(answer);
          break;
        }
        time_end = get_time();
        info.m = (time_end - time_start) / 60;
        info.s = (time_end - time_start) % 60;
        info.level = 3;
        if (judge(player_res, answer)) {
          printf("恭喜你成功了!\t用时:  %d:%d\n", info.m, info.s);
          record(info);
        } else {
          printf("回答错误!\t正确答案为:\n\n");
          fflush(stdin);
          print(answer);
        }
        break;
      default:
        pause("no option!");
        fflush(stdin);
        break;
      }
      break;
    case '2':
      easy = get_record(1); // 获取对应难度的记录
      normal = get_record(2);
      hard = get_record(3);
      order(easy); // 进行排序
      order(normal);
      order(hard);
      show(easy, normal, hard); // 显示排名
      pause("按任意键返回...");
      break;
    case '3':
      showHelp();
      pause("按任意键返回...");
      break;
    case '0':
      printf("\n拜拜~\n\n");
      exit(0);
    default:
      pause("输入有误！请重新输入...");
      break;
    }
  }
}

bool set(int x, int y, int val) {
  if (sudo[y][x] != 0) // 非空
    return false;
  int x0, y0;
  for (x0 = 0; x0 < 9; x0++) {
    if (sudo[y][x0] == val) // 行冲突
      return false;
  }
  for (y0 = 0; y0 < 9; y0++) {
    if (sudo[y0][x] == val) // 列冲突
      return false;
  }
  for (y0 = y / 3 * 3; y0 < y / 3 * 3 + 3; y0++) {
    for (x0 = x / 3 * 3; x0 < x / 3 * 3 + 3; x0++) {
      if (sudo[y0][x0] == val) // 格冲突
        return false;
    }
  }
  sudo[y][x] = val;
  return true;
}

void reset(int x, int y) { sudo[y][x] = 0; }

void initXOrd(int *xOrd) // 0~9随机序列
{
  int i, k, tmp;
  for (i = 0; i < 9; i++) {
    xOrd[i] = i;
  }
  for (i = 0; i < 9; i++) {
    k = rand() % 9;
    tmp = xOrd[k];
    xOrd[k] = xOrd[i];
    xOrd[i] = tmp;
  }
}

bool fillFrom(int y, int val) {
  int xOrd[9];
  initXOrd(xOrd); // 生成当前行的扫描序列
  for (int i = 0; i < 9; i++) {
    int x = xOrd[i];
    if (set(x, y, val)) {
      if (y == 8) // 到了最后一行
      {
        if (val == 9 ||
            fillFrom(0, val + 1)) // 当前填9则结束, 否则从第一行填下一个数
          return true;
      } else {
        if (fillFrom(y + 1, val)) // 下一行继续填当前数
          return true;
      }
      reset(x, y); // 回溯
    }
  }
  return false;
}

void initShudu() {
  srand(time(NULL));
  /*
          生成 9宫格
  */
  int i = 0, j = 0;
  for (i = 0; i < 9; i++) {
    for (j = 0; j < 9; j++) {
      sudo[i][j] = 0;
    }
  }
  while (!fillFrom(0, 1))
    ;
}

void get_answer(int *answer) {
  int i = 0, j = 0, k = 0;
  for (i = 0; i < MAXN; i++) {
    for (j = 0; j < MAXN; j++) {
      answer[k] = sudo[i][j];
      k++;
    }
  }
}

char printMainMenu() {
  char menuID;
  system("cls");
  printf("*************************************\n");
  printf("*            C语言数独游戏          *\n");
  printf("*************************************\n");
  printf("*          1.开始游戏               *\n");
  printf("*          2.查看排名               *\n");
  printf("*          3.玩法说明               *\n");
  printf("*          0.退出游戏               *\n");
  printf("*************************************\n");
  printf("*          请选择您的操作           *\n");
  printf("*************************************\n");
  fflush(stdin);
  menuID = getch();

  return menuID;
}

void showHelp() {
  printf("选择1 \"开始游戏\" 进入难度选择， 选择对应难度后即可进入游戏;\n");
  printf("选择2 \"查看排名\" 可以查看所有玩家的排名情况;\n");
  printf(
      "解答说明：解答需将数独完整写一遍，以空格分割每列，以回车分割每行!!!\n");
}

void print(int *answer) { // 打印数独
  printf("┏━━┳━━┳━━┳━━┳━━┳━━┳━━┳━━┳━━┓\n");
  for (int i = 0; i < 81; i++) {
    if (answer[i] == 0)
      printf("┃  ");
    else
      printf("┃ %d", answer[i]);
    if (i == 80) {
      printf("┃  ");
      printf("\n");
      printf("┗━━┻━━┻━━┻━━┻━━┻━━┻━━┻━━┻━━┛\n");
    } else if ((i + 1) % 9 == 0) {
      printf("┃  ");
      printf("\n");
      printf("┣━━╋━━╋━━╋━━╋━━╋━━╋━━╋━━╋━━┫\n");
    }
  }
}

void sudoku_level(int *answer, int count) { // 难度
  int x, y;                                 // 行号、列号
  int num = 0;
  srand(time(NULL));
  for (int i = 0; i < 81; i++)
    sudoku[i] = answer[i];
  while (num < (81 - count)) { // 挖空
    x = rand() % 9 + 1;
    y = rand() % 9 + 1;
    if (sudoku[(x - 1) * 9 + (y - 1)] != 0) {
      sudoku[(x - 1) * 9 + (y - 1)] = 0;
      num++;
    }
  }
}

bool receiver(int *player_res) { // 接收玩家答案
  for (int i = 0; i < 81; i++) {
    scanf("%d", &player_res[i]);
    if (!(player_res[i] >= 1 && player_res[i] <= 9)) { // 0 表示玩家放弃
      fflush(stdin);
      return false;
    }
  }
  return true;
}

int get_time() { // 获得当前时间秒
  time_t t;
  t = time(NULL);
  return t;
}

void ready() {
  print(sudoku);
  printf("你有5秒钟观察时间\n");
  for (int i = 0; i < 5; i++) {
    printf("●	");
    Sleep(1000);
  }
  printf("\n");
  printf("观察结束，计时开始，请开始作答。（输入除1~9外，视为放弃作答）\n");
  printf(
      "解答说明：解答需将数独完整写一遍，以空格分割每列，以回车分割每行!!!\n");
  printf("====================================================================="
         "=====\n");
}

bool judge(int *player_res, int *answer) { // 判断玩家答案
  for (int i = 0; i < 81; i++)
    if (player_res[i] != answer[i])
      return false;
  return true;
}

void record(player info) { // 记录
  FILE *fp;
  int M = MAX, S = MAX, LEVEL = MAX;
  char NAME[20];
  char remove[100] = {
      "                                                      "}; // 用于记录长度固定化，方便更新记录
  // 通过这种方法，可以直接在一个文件中更新数据，不必要全篇读—改—写，直接修改一行
  int c = 0;
  if ((fp = fopen("record.txt", "r+")) == NULL) { // 文件在cpp同目录下
    printf("文件不存在，保存失败！"); // 虽然会自动生成文件，but以防万一
    return;
  }
  setbuf(fp, NULL); // 设置缓冲区
  rewind(fp);
  c = ftell(fp); // 记录当前行的开头指针位置

  while (fscanf(fp, "%s %d:%d %d", NAME, &M, &S, &LEVEL) != EOF) {

    if (!strcmp(NAME, info.name) && LEVEL == info.level) { // strcmp比较相同返回0

      if (info.m < M || (info.m == M && info.s < S)) { // 如果是新纪录，则更新
        fseek(fp, c, SEEK_SET);
        fputs(remove, fp);      // 覆盖旧记录
        fseek(fp, c, SEEK_SET); // 回到该记录的开头位置
        fprintf(fp, "%s %d:%d %d", info.name, info.m, info.s,
                info.level); // 写入文件
        fflush(fp);          // 清除缓冲区
        return;
      }
      return; // 不是新纪录就不插入
    }
    fscanf(fp, "\n"); // 读取换行
    c = ftell(fp);
  }

  fputs(remove, fp);      // 先覆盖固定长度的区域
  fseek(fp, c, SEEK_SET); // 回到覆盖的区域首部
  fprintf(fp, "%s %d:%d %d", info.name, info.m, info.s,
          info.level);    // 在覆盖的区域内插入记录
  fseek(fp, 0, SEEK_END); // 指向尾部
  fprintf(fp, "\n");      // 插入换行符
  fclose(fp);
}

player *get_record(int level) { // 返回玩家记录的单向链表头结点
  FILE *fp;
  int M = MAX, S = MAX, LEVEL = MAX;
  char NAME[20];
  player *head = (player *)malloc(sizeof(player));
  head->next = NULL;
  if ((fp = fopen("record.txt", "r")) == NULL) {
    printf("文件不存在！");
    system("pause");
    exit(1);
  }
  setbuf(fp, NULL); // 设置缓冲区
  rewind(fp);
  while (fscanf(fp, "%s %d:%d %d", NAME, &M, &S, &LEVEL) != EOF) {
    if (LEVEL == level) {
      player *p = (player *)malloc(sizeof(player)); // 采用链表
      strcpy(p->name, NAME);
      p->m = M;
      p->s = S;
      p->next = head->next;
      head->next = p;
    }
  }
  fclose(fp);
  return head;
}

void order(player *head) { // 单链表排序
  player *p;
  player *q;
  int temp1;
  int temp2;
  char temp3[20];
  for (p = head->next; p != NULL; p = p->next)
    for (q = p->next; q != NULL; q = q->next)
      if (p->m > q->m || (p->m == q->m && p->s > q->s)) { // 对换两个结点的内容
        temp1 = p->m;
        temp2 = p->s;
        strcpy(temp3, p->name);
        p->m = q->m;
        p->s = q->s;
        strcpy(p->name, q->name);
        q->m = temp1;
        q->s = temp2;
        strcpy(q->name, temp3);
      }
}

void show(player *easy, player *normal, player *hard) { // 输出排行
  int no = 1;
  player *p1 = easy->next;
  player *p2 = normal->next;
  player *p3 = hard->next;
  printf("====================================================================="
         "=====\n");
  printf("\t\t   简单\t\t\t  一般\t\t\t   容易\n");
  while (p1 != NULL || p2 != NULL || p3 != NULL) {
    printf("NO.%d", no++);
    if (p1 != NULL) {
      printf("\t\t%s\t%d:%d\t", p1->name, p1->m, p1->s);
      p1 = p1->next;
    }
    if (p2 != NULL) {
      printf("\t%s\t%d:%d\t", p2->name, p2->m, p2->s);
      p2 = p2->next;
    }
    if (p3 != NULL) {
      printf("\t%s\t%d:%d\t", p3->name, p3->m, p3->s);
      p3 = p3->next;
    }
    printf("\n");
  }
}

// 暂停程序
void pause(const char *str, ...) {
  va_list vl;
  char buf[500] = {0};
  va_start(vl, str);
  vsnprintf(buf, 500, str, vl);
  va_end(vl);
  printf(buf);
  getch();
  printf("\n");
}

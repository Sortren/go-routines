from random import randint

PRODUCTS_NUMBER = 10_000


def main():
    with open("products.txt", "w") as f:
        for i in range(PRODUCTS_NUMBER):
            weight = randint(1, 500)
            f.write(f"{i} {weight}")


if __name__ == "__main__":
    main()

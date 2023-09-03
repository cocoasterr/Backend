# import time
# ### Read in the data with Pandas
# import pandas as pd

# s = time.time()
# df = pd.read_csv("train_split_00.csv")
# e = time.time()
# print("Pandas Loading Time = {}".format(e-s))

# ### Read in the data with Modin
# import ray
# ray.init(num_cpus=8)
# import modin.pandas as md_pd

# s = time.time()
# df = md_pd.read_csv("train_split_00.csv")
# e = time.time()
# print("Modin Loading Time = {}".format(e-s))

import csv
import time
import concurrent.futures
from concurrent.futures import ThreadPoolExecutor
import openpyxl
def process_csv_row(row, data):

    data.append(row)

def read_excel(filename):
    wb = openpyxl.load_workbook(filename)

    sheet = wb['Orders']
    data = []
    for row in sheet:
        process_csv_row(row, data)
    #using threadpool
    # with ThreadPoolExecutor(max_workers=16) as executor:
    #     for row in sheet.iter_rows(values_only=True):
    #         executor.submit(process_csv_row, row,data)
    return data
s = time.time()
data = read_excel("Orders-With Nulls.xlsx")
# Anda sekarang memiliki hasil pembacaan dalam results
e = time.time()
print("CSV Loading Time = {}".format(e - s))

#23.015033721923828
#18.320926427841187
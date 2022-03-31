from datetime import datetime

def getCurrentDate():
    return datetime.today().strftime("%m-%d-%y")

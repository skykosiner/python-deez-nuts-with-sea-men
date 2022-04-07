import datetime

from .User import Date

def get_current_date() -> Date:
    today = datetime.date.today()
    return Date(today.month, today.day, today.year)

def over_a_year(date: Date) -> bool:
    today = datetime.date.today()

    if date.month < today.month and date.day < today.day and date.year < today.year:
        return True
    return False

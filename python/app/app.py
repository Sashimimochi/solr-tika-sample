import tika
from tika import parser
import mysql.connector
from datetime import datetime

db_config = {
    'host': 'mysql',
    'user': 'mysql',
    'password': 'mysql',
    'database': 'book'
}
TABLE_NAME = 'documents'
FILEPATH = 'book.pdf'

tika.initVM()

def strpdate(date_string):
    return datetime.strptime(date_string, '%Y-%m-%dT%H:%M:%SZ')

def extract(filepath):
    return parser.from_file(filepath)

def create_table():
    query = f'''
CREATE TABLE IF NOT EXISTS {TABLE_NAME} (
    id INT AUTO_INCREMENT PRIMARY KEY,
    Title VARCHAR(255),
    Body TEXT,
    CreatedDate DATETIME,
    ModifiedDate DATETIME,
    Pages INT,
    CharLength INT
)
'''
    conn = mysql.connector.connect(**db_config)
    cursor = conn.cursor()

    try:
        cursor.execute(query)
        print("テーブルが作成されました")
    except mysql.connector.Error as err:
        print(f"テーブル作成中にエラーが発生しました: {err}")
    finally:
        cursor.close()
        conn.close()

def insert():
    parsed = extract(FILEPATH)
    data_to_insert = {
        'Title': parsed["metadata"]["dc:title"],
        'Body': parsed["content"],
        'CreatedDate': strpdate(parsed["metadata"]["dcterms:created"]),
        'ModifiedDate': strpdate(parsed["metadata"]["dcterms:modified"]),
        'Pages': parsed["metadata"]["xmpTPg:NPages"],
        'CharLength': sum([int(chars) for chars in parsed["metadata"]["pdf:charsPerPage"]])
    }

    query = f'''
INSERT INTO {TABLE_NAME} (Title, Body, CreatedDate, ModifiedDate, Pages, CharLength)
VALUES (%(Title)s, %(Body)s, %(CreatedDate)s, %(ModifiedDate)s, %(Pages)s, %(CharLength)s)
'''
    conn = mysql.connector.connect(**db_config)
    cursor = conn.cursor()

    try:
        cursor.execute(query, data_to_insert)
        conn.commit()
        print("データが挿入されました")
    except mysql.connector.Error as err:
        print(f"データの挿入中にエラーが発生しました: {err}")
    finally:
        cursor.close()
        conn.close()

def select():
    query = f'''
SELECT * FROM {TABLE_NAME}
'''

    # MySQLに接続し、カーソルを取得
    conn = mysql.connector.connect(**db_config)
    cursor = conn.cursor()

    try:
        # データを取得
        cursor.execute(query)
        rows = cursor.fetchall()

        # 取得したデータを表示
        print("挿入されたデータ:")
        for row in rows:
            print(row)
    except mysql.connector.Error as err:
        print(f"データの取得中にエラーが発生しました: {err}")
    finally:
        # 接続をクローズ
        cursor.close()
        conn.close()

def main():
    create_table()
    insert()
    select()

if __name__ == '__main__':
    main()

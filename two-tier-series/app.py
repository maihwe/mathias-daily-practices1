import os
import pymysql
pymysql.install_as_MySQLdb()
import pymysql

from flask import Flask, render_template, request, redirect, url_for
import pymysql

# Compatibility class to fake flask_mysqldb using pure PyMySQL
class MySQL:
    def __init__(self, app=None):
        self.app = app
        if app is not None:
            self.init_app(app)

    def init_app(self, app):
        pass

    @property
    def connection(self):
        # Dynamically grabs configurations directly from your existing app.config setup
        from flask import current_app
        return pymysql.connect(
            host=current_app.config.get('MYSQL_HOST', 'localhost'),
            user=current_app.config.get('MYSQL_USER', 'root'),
            password=current_app.config.get('MYSQL_PASSWORD', ''),
            database=current_app.config.get('MYSQL_DB', ''),
            cursorclass=pymysql.cursors.DictCursor # Ensures data comes back as dict if needed
        )
        
app = Flask(__name__)

# Configure MySQL from environment variables
app.config['MYSQL_HOST'] = os.environ.get('MYSQL_HOST', 'localhost')
app.config['MYSQL_USER'] = os.environ.get('MYSQL_USER', 'default_user')
app.config['MYSQL_PASSWORD'] = os.environ.get('MYSQL_PASSWORD', 'default_password')
app.config['MYSQL_DB'] = os.environ.get('MYSQL_DB', 'default_db')

# Initialize MySQL
mysql = MySQL(app)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/register', methods=['POST'])
def register():
    name = request.form.get('name')
    phone_number = request.form.get('phonenumber')
    email = request.form.get('email')
    country = request.form.get('country')
    
    # Insert registration data into MySQL
    cur = mysql.connection.cursor()
    cur.execute('INSERT INTO registrations (name, phone_number, email, country) VALUES (%s, %s, %s, %s)', 
                (name, phone_number, email, country))
    mysql.connection.commit()
    cur.close()
    
    return redirect(url_for('index'))

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
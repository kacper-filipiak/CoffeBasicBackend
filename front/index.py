from flask import Flask, request, render_template, url_for, redirect, render_template, session
import datetime
from requests import Session
import requests

app = Flask(__name__)
app.secret_key = b'12372329874'

def get_session():
    session = Session()
    session.trust_env = False
    return session

navigation = {'Home':'/', 'Add coffe type':'/type/add', 'Show coffe types':'/type/', 'Add coffe bag':'/bag/add', 'Show coffe bags':'/bag/'}

@app.route(navigation['Home'])
def hello_world():
    text = session.get('text', "")
    body = get_session().get('http://api:8080/coffe/types')
    return render_template('index.html', navigation=navigation, body=body.text)

@app.get('/type/add')
def add_type_form():
    return render_template('add-form.html', navigation=navigation)

@app.post('/type/add')
def add_type_form_post():
    type_name = request.form['type_name']
    body = get_session().post('http://api:8080/coffe/types', json={'Name':type_name})
    print(body)
    return redirect('/', 200)

@app.get('/type/')
def show_type():
    types = get_session().get('http://api:8080/coffe/types').json()['coffeTypes'] 
    return render_template('type-list.html', navigation=navigation, types=types)



@app.get('/bag/')
def show_bag():
    bags = get_session().get('http://api:8080/coffe/bags').json()['coffeBags']
    for bag in bags: 
        bag.update({'date': datetime.datetime.fromtimestamp(int(bag['date'])).strftime('%d/%m/%Y')})
    return render_template('bag-list.html', navigation=navigation, bags=bags)


@app.get('/bag/add')
def add_bag_form():
    types = get_session().get('http://api:8080/coffe/types').json()['coffeTypes'] 
    return render_template('add-bag-form.html', navigation=navigation, types=types)

@app.post('/bag/add')
def add_bag_form_post():
    type_name = request.form['type_name']
    print(type_name)
    weight = request.form['weight']
    roast_date = request.form['roast_date']
    body = get_session().post('http://api:8080/coffe/bags',
                         json={
                             'TypeName':type_name, 
                             'Weight': float(weight), 
                             'RoastDate':roast_date
                             })
    print(body.json())
    return redirect('/', 200)

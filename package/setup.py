from setuptools import setup, find_packages

setup(
    name='mypackage',
    version='1.0.2',
    description='mypackage description',
    author='koketani',
    author_email='yyyyyykoketani@gmail.com',
    python_requires='>=3.6.0',
    url='https://github.com/koketani/study-python/pakcage/myproject',
    packages=find_packages(exclude=('tests',)),
    install_requires=['requests'],
    extras_require={
        'play': 'jinja2',
    },
    include_package_data=True,
)

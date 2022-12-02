import axios from 'axios';
import React from 'react';
import User from '../types/type'

export function getUsers ()
{
     axios.get('http://localhost:8080/users/')
    .then(function (response) {
        // handle success
        console.log(response);
    })
    .catch(function (error) {
        // handle error
        console.log(error);
    })
    .finally(function () {
        // always executed
    });
}


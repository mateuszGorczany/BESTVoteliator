import { useState, useEffect } from "react";
import { clientID } from "../utils";
// import {GoogleButton from "./GoogleButton"

function Navbar() {
    return (
        <nav className="bg-white border-gray-200 px-2 sm:px-4 py-2.5 rounded dark:bg-gray-900">
            <div className="container flex flex-wrap justify-between items-center mx-auto">
                <a href="https://www.newsite.best.krakow.pl/" className="flex items-center">
                    <img src="https://www.newsite.best.krakow.pl/themes/img/logo.svg" className="mr-3 h-6 sm:h-9" alt="BEST Logo" />
                </a>
                <a href="/">
                    <span className="self-center text-xl font-semibold whitespace-nowrap dark:text-white">BESTVoteliator</span>
                </a>
                <div id="signInDiv" className="flex md:order-2">
                    <div id="signInDiv"></div>
                </div>
                <div className="hidden justify-between items-center w-full md:flex md:w-auto md:order-1" id="navbar-cta">
                    <ul className="flex flex-col p-4 mt-4 bg-gray-50 rounded-lg border border-gray-100 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
                        <li>
                            <a href="/#" className="block py-2 pr-4 pl-3 text-white bg-blue-700 rounded md:bg-transparent md:text-blue-700 md:p-0 dark:text-white" aria-current="page">Strona Główna</a>
                        </li>
                        <li>
                            <a href="/#/create" className="block py-2 pr-4 pl-3 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0 md:dark:hover:text-white dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">Stwórz głosowanie</a>
                        </li>
                        <li>
                            <a href="/#/about" className="block py-2 pr-4 pl-3 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-blue-700 md:p-0 md:dark:hover:text-white dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">O serwisie</a>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
}

export default Navbar
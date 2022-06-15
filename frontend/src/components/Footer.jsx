import React from 'react'
import logoImg from '../assets/logohi.png'
import {MailIcon, HomeIcon, PhoneIcon} from '@heroicons/react/outline'

const Footer = () => {
  return (
    <div className='w-full mt-24 bg-black text-white py-y px-2'>
        <div className='max-w-[1240px] mx-auto grid grid-cols-2 md:grid-cols-6 border-b-2 border-gray-600 py-8'>
            <div className='col-span-2 px-3'>
                <img className='bg-white h-10 w-18' src={logoImg} alt="Hicoder" /><br/>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Deserunt itaque perspiciatis voluptas dicta, aperiam quae expedita consectetur assumenda unde quos quis odio ex quod voluptate veritatis dolorem odit nemo nobis.</p><br/>
            </div>
            <div className='col-span-1 px-8'>
                <h6 className='font-bold uppercase pt-2'>Content</h6>
                <ul>
                    <li className='py-1'><a href='/'>Home</a></li>
                    <li className='py-1'><a href='/'>About</a></li>
                    <li className='py-1'><a href='/'>Find Mentor</a></li>
                </ul>
            </div>
            <div className='col-span-1 px-8'>
                <h6 className='font-bold uppercase pt-2'>Legal</h6>
                <ul>
                    <li className='py-1'>Claims</li>
                    <li className='py-1'>Privacy</li>
                    <li className='py-1'>Terms</li>
                    <li className='py-1'>Policies</li>
                    <li className='py-1'>Conditions</li>
                </ul>
            </div>
            <div className='col-span-2 pt-8 md:pt-2 px-3'>
                <p className='font-bold uppercase'>Contact Us</p><br/>
                <form className='flex flex-col sm:flex-row'>
                    <input className='w-full p-2 mr-4 rounded-md mb-4' type="email" placeholder='Enter email..'/>
                    <button className='p-2 mb-4'>Subscribe</button>
                </form>
                <div className='flex flex-row'>
                    <MailIcon className='w-8 mx-3'/>
                    <p>engineering11@ruangguru.com</p>
                </div>
                <div className='flex flex-row'>
                    <HomeIcon className='w-8 mx-3'/>
                    <p>engineering11, Ruangguru, Indonesia</p>
                </div>
                <div className='flex flex-row'>
                    <PhoneIcon className='w-8 mx-3'/>
                    <p>+62 1234-5678</p>
                </div>
            </div>
        </div>

        <div className='max-w-[1240px] px-2 py-4 mx-auto text-center text-white'>
            <p className='py-4'>Copyright 2022. All rights reserved</p>
        </div>
    </div>
  )
}

export default Footer
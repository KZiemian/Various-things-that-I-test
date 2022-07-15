#!/usr/bin/python3
# -*- coding: utf-8 -*-

do_oddania = 5550.0

oddales = []



####################
# 2015-2016
oddales.append(100.0)
oddales.append(200.0)
oddales.append(300.0)
oddales.append(400.0)
oddales.append(152.0)
oddales.append(292.0)
oddales.append(132.0)
oddales.append(132.0)
oddales.append(143.0)
oddales.append(150.0)
oddales.append(250.0)
oddales.append(500.0)



####################
# 2017 rok
oddales.append(500.0)
oddales.append(250.0)
oddales.append(300.0)  # 25 XI
oddales.append(200.0)  # 4-5 XII
oddales.append(200.0)  # 15 XII



####################
# 2018 rok

# Wielki Post
oddales.append(200.0)  # 15 II
# Tu się skończył Wielki Post. Jesteś żałosny

oddales.append(183.20)
oddales.append(100)  # 1 VII
oddales.append(200)  # 16 VIII
oddales.append(65)  # 27 VIII
oddales.append(100)  # 12 IX
oddales.append(100)  # 18 IX
oddales.append(100)  # 12 XI
oddales.append(200)  # 22 XI
oddales.append(110)  # 24 XII

# Zajęło ci to tyle czasu, żałosny jesteś.


print("Oddales juz: {}.".format(round(sum(oddales, 0.0))))
zostalo_do_oddania = round(do_oddania - sum(oddales, 0.0), 2)
print("Do oddania pozostalo: {}.".format(zostalo_do_oddania))
